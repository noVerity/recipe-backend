package rest

import (
	"adomeit.xyz/recipe/core"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, auth *AuthManager, recipeCore *core.RecipeCore, ingredientCore *core.IngredientCore) *gin.Engine {
	NewIngredientController(r, ingredientCore, auth)
	NewRecipeController(r, recipeCore, auth)
	return r
}
