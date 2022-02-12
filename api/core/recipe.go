package core

import (
	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/ent/recipe"
	"adomeit.xyz/recipe/mq"
	"context"
	"github.com/gosimple/slug"
	"github.com/noVerity/gofromto"
	"sort"
)

// Recipe is a model used to exchange information with the frontend
type Recipe struct {
	Id              string
	Slug            string
	Name            string
	IngredientsList string
	Instructions    string
	Nutrition       string
	Servings        int
	User            string
}

// RecipeCore provides CRUD operations on our recipe model store
type RecipeCore struct {
	client *ent.Client
	queue  *mq.MQ
}

// NewRecipeCore creates a new instance
func NewRecipeCore(client *ent.Client, queue *mq.MQ) *RecipeCore {
	return &RecipeCore{client, queue}
}

// Create tries to write a new recipe into the DB, will request info on unknown ingredients
func (core *RecipeCore) Create(newRecipe Recipe, c context.Context) (*ent.Recipe, error) {
	recipeSlug := slug.Make(newRecipe.Name)

	foundIngredients, missingIngredients := core.getIngredientsFromList(newRecipe.IngredientsList, c)

	// TODO: Calculate nutritional value for the recipe

	createdRecipe, err := core.client.Recipe.Create().
		SetName(newRecipe.Name).
		SetSlug(recipeSlug).
		SetIngredientslist(newRecipe.IngredientsList).
		SetInstructions(newRecipe.Instructions).
		SetNutrition("").
		SetServings(newRecipe.Servings).
		SetUser(newRecipe.User).
		AddIngredients(foundIngredients...).
		Save(c)

	if err != nil {
		return &ent.Recipe{}, err
	}

	if core.queue != nil {
		go core.queue.RequestIngredients(missingIngredients, createdRecipe.ID)
	}

	return createdRecipe, err
}

// Update an existing recipe entry
func (core *RecipeCore) Update(id string, update Recipe, c context.Context) (*ent.Recipe, error) {
	previous, err := core.client.Recipe.Query().
		Where(recipe.ID(id)).
		First(c)

	if err != nil {
		return &ent.Recipe{}, err
	}

	foundIngredients, missingIngredients := core.getIngredientsFromList(update.IngredientsList, c)

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

	result, err := core.client.Recipe.UpdateOne(previous).
		SetName(update.Name).
		SetSlug(slug.Make(update.Name)).
		SetInstructions(update.Instructions).
		SetServings(update.Servings).
		SetIngredientslist(update.IngredientsList).
		RemoveIngredients(ingredientsToRemove...).
		AddIngredients(ingredientsToAdd...).
		Save(c)

	if err != nil {
		return &ent.Recipe{}, err
	}

	if core.queue != nil {
		go core.queue.RequestIngredients(missingIngredients, result.ID)
	}

	return result, nil
}

// Get the Recipe with the given ID
func (core *RecipeCore) Get(id string, c context.Context) (*ent.Recipe, error) {
	return core.client.Recipe.Query().
		Where(recipe.ID(id)).
		WithIngredients().
		First(c)
}

// FindAll recipes for the given user, returns the total count and the recipes that fall in the given limit and offset
func (core *RecipeCore) FindAll(user string, limit int, offset int, c context.Context) (int, []*ent.Recipe, error) {
	if limit == 0 || limit > 1000 {
		limit = 1000
	}

	count, err := core.client.
		Recipe.
		Query().
		Where(recipe.User(user)).
		Count(c)

	if err != nil {
		return 0, make([]*ent.Recipe, 0), err
	}

	if count == 0 || count < offset {
		return 0, make([]*ent.Recipe, 0), nil
	}

	results, err := core.client.Recipe.Query().
		Limit(limit).
		Offset(offset).
		All(c)

	return count, results, err
}

// Delete removes the recipe from the DB
func (core *RecipeCore) Delete(id string, c context.Context) (*ent.Recipe, error) {
	result, err := core.client.Recipe.Query().
		Where(recipe.ID(id)).
		First(c)

	if err != nil {
		return &ent.Recipe{}, err
	}

	err = core.client.Recipe.DeleteOne(result).Exec(c)
	return result, err
}

// getIngredientsFromList parses a string list for ingredients and checks them against what we have saved in the DB
func (core *RecipeCore) getIngredientsFromList(list string, c context.Context) ([]*ent.Ingredient, []gofromto.Measure) {
	ingredients := ParseIngredientList(list)

	ingredientNames := make([]string, len(ingredients))

	for i, entry := range ingredients {
		ingredientNames[i] = entry.Name
	}

	foundIngredients, _ := core.client.Ingredient.
		Query().
		Where(ingredient.NameIn(ingredientNames...)).
		All(c)

	var missingIngredients []gofromto.Measure

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
