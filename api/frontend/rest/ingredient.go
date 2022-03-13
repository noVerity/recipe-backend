package rest

import (
	"adomeit.xyz/recipe/core"
	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/shared"
	"github.com/gin-gonic/gin"
	"net/http"
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

func ingredientModelToResponse(modelEntity *ent.Ingredient) Ingredient {
	return Ingredient{
		Name:          modelEntity.Name,
		Calories:      modelEntity.Calories,
		Fat:           modelEntity.Fat,
		Carbohydrates: modelEntity.Carbohydrates,
		Protein:       modelEntity.Protein,
	}
}

func ingredientsToResponse(modelEntities []*ent.Ingredient) []Ingredient {
	responseSlice := make([]Ingredient, len(modelEntities))
	for i, element := range modelEntities {
		responseSlice[i] = ingredientModelToResponse(element)
	}
	return responseSlice
}

func (ingredient Ingredient) ToModel() core.Ingredient {
	return core.Ingredient{
		Name:          ingredient.Name,
		Calories:      ingredient.Calories,
		Fat:           ingredient.Fat,
		Carbohydrates: ingredient.Carbohydrates,
		Protein:       ingredient.Protein,
	}
}

type IngredientController struct {
	core *core.IngredientCore
}

// NewIngredientController takes the gin engine and creates routes for CRUD on ingredients
func NewIngredientController(r *gin.Engine, core *core.IngredientCore, auth *shared.AuthManager) *IngredientController {
	controller := IngredientController{core}
	userRoute := r.Group("/ingredient", auth.AuthMiddleware())
	{
		userRoute.POST("", controller.HandleCreateIngredient)
		userRoute.GET("", controller.HandleGetAllIngredients)
		userRoute.PUT("/:id", controller.HandleUpdateIngredient)
		userRoute.PATCH("/:id", controller.HandleUpdateIngredient)
		userRoute.GET("/:id", controller.HandleGetIngredient)
		userRoute.DELETE("/:id", controller.HandleDeleteIngredient)
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

	createdIngredient, err := controller.core.Create(newIngredient.ToModel(), c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ingredientModelToResponse(createdIngredient))
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

	result, err := controller.core.Update(uriElement.Id, update.ToModel(), c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ingredientModelToResponse(result))
}

// HandleGetIngredient returns an ingredient
func (controller *IngredientController) HandleGetIngredient(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.core.Get(uriElement.Id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ingredient not found"})
		return
	}

	c.JSON(http.StatusOK, ingredientModelToResponse(result))
}

// HandleGetAllIngredients returns all ingredients within offset and limit
func (controller *IngredientController) HandleGetAllIngredients(c *gin.Context) {
	var query QueryPagination
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Limit == 0 || query.Limit > 1000 {
		query.Limit = 1000
	}

	count, result, err := controller.core.GetAll(query.Limit, query.Offset, c)

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

	c.JSON(http.StatusOK, IngredientsResult{
		PagedResult{
			Pagination: PagingInformation{
				Offset: query.Offset,
				Count:  count,
			},
		},
		ingredientsToResponse(result),
	})
}

// HandleDeleteIngredient removes an ingredient
func (controller *IngredientController) HandleDeleteIngredient(c *gin.Context) {
	var uriElement URIElement
	if err := c.BindUri(&uriElement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.core.Delete(uriElement.Id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ingredient not found"})
		return
	}

	c.JSON(http.StatusOK, ingredientModelToResponse(result))
}
