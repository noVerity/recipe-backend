package rest

import (
	"adomeit.xyz/recipe/core"
	"adomeit.xyz/recipe/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Recipe struct {
	Id              string       `json:"id"`
	Slug            string       `json:"slug"`
	Name            string       `json:"name"`
	IngredientsList string       `json:"ingredientslist"`
	Instructions    string       `json:"instructions"`
	Nutrition       string       `json:"nutrition"`
	Servings        int          `json:"servings"`
	Ingredients     []Ingredient `json:"ingredients,omitempty"`
	User            string       `json:"user"`
}

type RecipeResult struct {
	PagedResult
	Data []Recipe `json:"data"`
}

func recipeModelToResponse(modelEntity *ent.Recipe) Recipe {
	var ingredients []Ingredient
	for _, entry := range modelEntity.Edges.Ingredients {
		ingredients = append(ingredients, ingredientModelToResponse(entry))
	}
	return Recipe{
		Id:              modelEntity.ID,
		Slug:            modelEntity.Slug,
		Name:            modelEntity.Name,
		IngredientsList: modelEntity.Ingredientslist,
		Instructions:    modelEntity.Instructions,
		Nutrition:       modelEntity.Nutrition,
		Servings:        modelEntity.Servings,
		User:            modelEntity.User,
		Ingredients:     ingredients,
	}
}

func (recipe Recipe) ToModel() core.Recipe {
	return core.Recipe{
		Name:            recipe.Name,
		Slug:            recipe.Slug,
		IngredientsList: recipe.IngredientsList,
		Instructions:    recipe.Instructions,
		Nutrition:       recipe.Nutrition,
		Servings:        recipe.Servings,
		User:            recipe.User,
	}
}

func recipesToResponse(modelEntities []*ent.Recipe) []Recipe {
	responseSlice := make([]Recipe, len(modelEntities))
	for i, element := range modelEntities {
		responseSlice[i] = recipeModelToResponse(element)
	}
	return responseSlice
}

type RecipeController struct {
	core *core.RecipeCore
}

// NewRecipeController takes the gin engine and creates routes for CRUD on recipes
func NewRecipeController(r *gin.Engine, core *core.RecipeCore, auth *AuthManager) *RecipeController {
	controller := RecipeController{core}
	userRoute := r.Group("/recipe", auth.AuthMiddleware())
	{
		userRoute.POST("", controller.HandleCreateRecipe)
		userRoute.GET("/:id", controller.HandleGetRecipe)
		userRoute.PUT("/:id", controller.HandleUpdateRecipe)
		userRoute.PATCH("/:id", controller.HandleUpdateRecipe)
		userRoute.GET("", controller.HandleGetAllRecipes)
		userRoute.DELETE("/:id", controller.HandleDeleteRecipe)
	}
	return &controller
}

func (controller *RecipeController) HandleCreateRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.BindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecipe.User = c.GetString("user")

	createdRecipe, err := controller.core.Create(newRecipe.ToModel(), c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "could not create recipe"})
		return
	}

	c.JSON(http.StatusCreated, recipeModelToResponse(createdRecipe))
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

	result, err := controller.core.Update(uriElement.Id, update.ToModel(), c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "unable to update recipe"})
		return
	}

	c.JSON(http.StatusOK, recipeModelToResponse(result))
}

func (controller *RecipeController) HandleGetRecipe(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.core.Get(uriElement.Id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipeModelToResponse(result))
}

func (controller *RecipeController) HandleGetAllRecipes(c *gin.Context) {
	var query QueryPagination
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, result, err := controller.core.FindAll(c.GetString("user"), query.Limit, query.Offset, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "encountered error"})
		return
	}

	c.JSON(http.StatusOK, RecipeResult{
		PagedResult{
			Pagination: PagingInformation{
				Offset: query.Offset,
				Count:  count,
			},
		},
		recipesToResponse(result),
	})
}

// HandleDeleteRecipe removes a recipe
func (controller *RecipeController) HandleDeleteRecipe(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.core.Delete(uriElement.Id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipeModelToResponse(result))
}
