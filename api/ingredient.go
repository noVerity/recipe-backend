package main

import (
	"net/http"

	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/ingredient"

	"github.com/gin-gonic/gin"
)

type Ingredient struct {
	Name          string  `json:"name"`
	Calories      float32 `json:"calories"`
	Fat           float32 `json:"fat"`
	Carbohydrates float32 `json:"carbohydrates"`
	Protein       float32 `json:"protein"`
}

type IngredientsResult struct {
	PagedResult
	Data []Ingredient `json:"data"`
}

func IngredientModelToResponse(modelEntity *ent.Ingredient) Ingredient {
	return Ingredient{
		Name:          modelEntity.Name,
		Calories:      modelEntity.Calories,
		Fat:           modelEntity.Fat,
		Carbohydrates: modelEntity.Carbohydrates,
		Protein:       modelEntity.Protein,
	}
}

func IngredientsToResponse(modelEntities []*ent.Ingredient) []Ingredient {
	responseSlice := make([]Ingredient, len(modelEntities))
	for i, element := range modelEntities {
		responseSlice[i] = IngredientModelToResponse(element)
	}
	return responseSlice
}

type IngredientController struct {
	router *gin.Engine
	client *ent.Client
}

// NewIngredientController takes the gin engine and creates routes for CRUD on ingredients
func NewIngredientController(r *gin.Engine, client *ent.Client) *IngredientController {
	controller := IngredientController{r, client}
	userRoute := r.Group("/ingredient")
	{
		userRoute.POST("", controller.HandleCreateIngredient)
		userRoute.GET("", controller.HandleGetAllIngredients)
		userRoute.PUT("/:name", controller.HandleUpdateIngredient)
		userRoute.PATCH("/:name", controller.HandleUpdateIngredient)
		userRoute.GET("/:name", controller.HandleGetIngredient)
		userRoute.DELETE("/:name", controller.HandleDeleteIngredient)
	}
	return &controller
}

// HandleCreateIngredient creates a new ingredient
func (controller *IngredientController) HandleCreateIngredient(c *gin.Context) {
	var newIngredient Ingredient
	if err := c.BindJSON(&newIngredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := controller.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(newIngredient.Name)).
		First(c)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "ingredient with this name already exists"})
		return
	}

	createdIngredient, err := controller.client.Ingredient.Create().
		SetName(newIngredient.Name).
		SetCalories(newIngredient.Calories).
		SetFat(newIngredient.Fat).
		SetCarbohydrates(newIngredient.Carbohydrates).
		SetProtein(newIngredient.Protein).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "ingredient with this name already exists"})
		return
	}

	c.JSON(http.StatusCreated, IngredientModelToResponse(createdIngredient))
}

// HandleUpdateIngredient updates an ingredient
func (controller *IngredientController) HandleUpdateIngredient(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var update Ingredient
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	previous, err := controller.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(uriElement.Name)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ingredient does not exist"})
		return
	}

	result, err := controller.client.Ingredient.UpdateOne(previous).
		SetName(update.Name).
		SetCalories(update.Calories).
		SetFat(update.Fat).
		SetCarbohydrates(update.Carbohydrates).
		SetProtein(update.Protein).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, IngredientModelToResponse(result))
}

// HandleGetIngredient returns an ingredient
func (controller *IngredientController) HandleGetIngredient(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(uriElement.Name)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ingredient not found"})
		return
	}

	c.JSON(http.StatusOK, IngredientModelToResponse(result))
}

// HandleGetIngredient returns an ingredient
func (controller *IngredientController) HandleGetAllIngredients(c *gin.Context) {
	var query QueryPagination
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Limit == 0 || query.Limit > 1000 {
		query.Limit = 1000
	}

	count, err := controller.client.
		Ingredient.
		Query().
		Count(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "encountered error"})
		return
	}

	if count == 0 || count < query.Offset {
		c.JSON(http.StatusOK, IngredientsResult{
			PagedResult{
				Pagination: PagingInformation{
					Offset: query.Offset,
					Count:  count,
				},
			},
			make([]Ingredient, 0),
		})
		return
	}

	result, err := controller.client.Ingredient.Query().
		Limit(query.Limit).
		Offset(query.Offset).
		All(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "nothing found"})
		return
	}

	c.JSON(http.StatusOK, IngredientsResult{
		PagedResult{
			Pagination: PagingInformation{
				Offset: query.Offset,
				Count:  count,
			},
		},
		IngredientsToResponse(result),
	})
}

// HandleDeleteIngredient removes an ingredient
func (controller *IngredientController) HandleDeleteIngredient(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(uriElement.Name)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ingredient not found"})
		return
	}

	err = controller.client.Ingredient.DeleteOne(result).Exec(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, IngredientModelToResponse(result))
}
