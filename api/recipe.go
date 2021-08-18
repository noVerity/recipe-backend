package main

import (
	"net/http"
	"sort"

	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/ent/recipe"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type Recipe struct {
	Name            string       `json:"name"`
	IngredientsList string       `json:"ingredientslist"`
	Instructions    string       `json:"instructions"`
	Nutrition       string       `json:"nutrition"`
	Servings        int          `json:"servings"`
	Ingredients     []Ingredient `json:"ingredients,omitempty"`
}

type RecipeResult struct {
	PagedResult
	Data []Recipe `json:"data"`
}

func RecipeModelToResponse(modelEntity *ent.Recipe) Recipe {
	var ingredients []Ingredient
	for _, entry := range modelEntity.Edges.Ingredients {
		ingredients = append(ingredients, IngredientModelToResponse(entry))
	}
	return Recipe{
		Name:            modelEntity.Name,
		IngredientsList: modelEntity.Ingredientslist,
		Instructions:    modelEntity.Instructions,
		Nutrition:       modelEntity.Nutrition,
		Servings:        modelEntity.Servings,
		Ingredients:     ingredients,
	}
}

func RecipesToResponse(modelEntities []*ent.Recipe) []Recipe {
	responseSlice := make([]Recipe, len(modelEntities))
	for i, element := range modelEntities {
		responseSlice[i] = RecipeModelToResponse(element)
	}
	return responseSlice
}

type RecipeController struct {
	router             *gin.Engine
	client             *ent.Client
	requestIngredients func(ingredients []IngredientEntry, recipeId int)
}

// NewRecipeController takes the gin engine and creates routes for CRUD on recipes
func NewRecipeController(r *gin.Engine, client *ent.Client, auth *AuthManager, requestIngredients func(ingredients []IngredientEntry, recipeId int)) *RecipeController {
	controller := RecipeController{r, client, requestIngredients}
	userRoute := r.Group("/recipe", auth.AuthMiddleware())
	{
		userRoute.POST("", controller.HandleCreateRecipe)
		userRoute.GET("/:name", controller.HandleGetRecipe)
		userRoute.PUT("/:name", controller.HandleUpdateRecipe)
		userRoute.PATCH("/:name", controller.HandleUpdateRecipe)
		userRoute.GET("", controller.HandleGetAllRecipes)
		userRoute.DELETE("/:name", controller.HandleDeleteRecipe)
	}
	return &controller
}

func (controller *RecipeController) HandleCreateRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.BindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slug := slug.Make(newRecipe.Name)

	_, err := controller.client.Recipe.Query().
		Where(recipe.Slug(slug)).
		First(c)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "recipe with this name already exists"})
		return
	}

	foundIngredients, missingIngredients := controller.GetIngredientsFromList(newRecipe.IngredientsList, c)

	// TODO: Calculate nutritional value for the recipe

	createdRecipe, err := controller.client.Recipe.Create().
		SetName(newRecipe.Name).
		SetSlug(slug).
		SetIngredientslist(newRecipe.IngredientsList).
		SetInstructions(newRecipe.Instructions).
		SetNutrition("").
		SetServings(newRecipe.Servings).
		AddIngredients(foundIngredients...).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "recipe with this name already exists"})
		return
	}

	controller.requestIngredients(missingIngredients, createdRecipe.ID)

	c.JSON(http.StatusCreated, RecipeModelToResponse(createdRecipe))
}

func (controller *RecipeController) HandleUpdateRecipe(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var update Recipe
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	previous, err := controller.client.Recipe.Query().
		Where(recipe.Slug(uriElement.Name)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe does not exist"})
		return
	}

	foundIngredients, missingIngredients := controller.GetIngredientsFromList(update.IngredientsList, c)

	var ingredientsToRemove []*ent.Ingredient
	var ingredientsToAdd []*ent.Ingredient

	for i := range previous.Edges.Ingredients {
		index := sort.Search(len(foundIngredients), func(n int) bool {
			return previous.Edges.Ingredients[i].ID == foundIngredients[n].ID
		})
		if index < 0 {
			ingredientsToRemove = append(ingredientsToRemove, previous.Edges.Ingredients[i])
		}
	}

	for i := range foundIngredients {
		index := sort.Search(len(previous.Edges.Ingredients), func(n int) bool {
			return previous.Edges.Ingredients[n].ID == foundIngredients[i].ID
		})
		if index < 0 {
			ingredientsToAdd = append(ingredientsToAdd, foundIngredients[i])
		}
	}

	result, err := controller.client.Recipe.UpdateOne(previous).
		SetName(update.Name).
		SetInstructions(update.Instructions).
		SetServings(update.Servings).
		SetIngredientslist(update.IngredientsList).
		RemoveIngredients(ingredientsToRemove...).
		AddIngredients(ingredientsToAdd...).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	controller.requestIngredients(missingIngredients, result.ID)

	c.JSON(http.StatusOK, RecipeModelToResponse(result))
}

func (controller *RecipeController) HandleGetRecipe(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.client.Recipe.Query().
		Where(recipe.SlugEqualFold(uriElement.Name)).
		WithIngredients().
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	c.JSON(http.StatusOK, RecipeModelToResponse(result))
}

func (controller *RecipeController) HandleGetAllRecipes(c *gin.Context) {
	var query QueryPagination
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Limit == 0 || query.Limit > 1000 {
		query.Limit = 1000
	}

	count, err := controller.client.
		Recipe.
		Query().
		Count(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "encountered error"})
		return
	}

	if count == 0 || count < query.Offset {
		c.JSON(http.StatusOK, RecipeResult{
			PagedResult{
				Pagination: PagingInformation{
					Offset: query.Offset,
					Count:  count,
				},
			},
			make([]Recipe, 0),
		})
		return
	}

	result, err := controller.client.Recipe.Query().
		Limit(query.Limit).
		Offset(query.Offset).
		All(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "nothing found"})
		return
	}

	c.JSON(http.StatusOK, RecipeResult{
		PagedResult{
			Pagination: PagingInformation{
				Offset: query.Offset,
				Count:  count,
			},
		},
		RecipesToResponse(result),
	})
}

// HandleDeleteRecipe removes a recipe
func (controller *RecipeController) HandleDeleteRecipe(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.client.Recipe.Query().
		Where(recipe.Slug(uriElement.Name)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	err = controller.client.Recipe.DeleteOne(result).Exec(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RecipeModelToResponse(result))
}

func (controller *RecipeController) GetIngredientsFromList(list string, c *gin.Context) ([]*ent.Ingredient, []IngredientEntry) {
	ingredients := ParseIngredientList(list)

	ingredientNames := make([]string, len(ingredients))

	for i, entry := range ingredients {
		ingredientNames[i] = entry.Name
	}

	foundIngredients, _ := controller.client.Ingredient.
		Query().
		Where(ingredient.NameIn(ingredientNames...)).
		All(c)

	var missingIngredients []IngredientEntry

	for _, entry := range ingredients {
		found := false
		for _, foundEntry := range foundIngredients {
			if foundEntry.Name == entry.Name {
				found = true
				continue
			}
		}
		if !found {
			missingIngredients = append(missingIngredients, entry)
		}
	}

	return foundIngredients, missingIngredients
}
