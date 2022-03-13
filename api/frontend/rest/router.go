package rest

import (
	"adomeit.xyz/recipe/core"
	"adomeit.xyz/shared"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, auth *shared.AuthManager, recipeCore *core.RecipeCore, ingredientCore *core.IngredientCore) *gin.Engine {
	r.Use(ThrottleMiddleware())
	NewIngredientController(r, ingredientCore, auth)
	NewRecipeController(r, recipeCore, auth)
	return r
}
