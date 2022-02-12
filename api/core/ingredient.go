package core

import (
	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/mq"
	"context"
	"fmt"
)

// Ingredient is a model used to exchange information with the frontend
type Ingredient struct {
	Name          string  `json:"name"`
	Calories      float32 `json:"calories"`
	Fat           float32 `json:"fat"`
	Carbohydrates float32 `json:"carbohydrates"`
	Protein       float32 `json:"protein"`
}

// IngredientCore handles the business logic around ingredients
type IngredientCore struct {
	client *ent.Client
	queue  *mq.MQ
}

// NewIngredientCore creates a new ingredient core instance
func NewIngredientCore(client *ent.Client, queue *mq.MQ) *IngredientCore {
	return &IngredientCore{client, queue}
}

// AcceptIngredientResults subscribes to the result messages from the message queue and creates ingredients from them
func (core *IngredientCore) AcceptIngredientResults(c context.Context) error {
	results, err := core.queue.AcceptIngredientResults()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-c.Done():
				fmt.Println("Closing accept ingredients process")
				return
			case result := <-results:
				created, err := core.Create(Ingredient{
					result.Name,
					result.Calories,
					result.Fat,
					result.Carbohydrates,
					result.Protein,
				}, c)
				if err != nil {
					fmt.Printf("Failed to create ingredient: %v\n", err)
				} else {
					fmt.Printf("Created new ingredient from queue: %v\n", created.Name)
				}
			}
		}
	}()

	return nil
}

// Create a new ingredient
func (core *IngredientCore) Create(newIngredient Ingredient, c context.Context) (*ent.Ingredient, error) {
	_, err := core.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(newIngredient.Name)).
		First(c)

	if err == nil {
		return &ent.Ingredient{}, fmt.Errorf("ingredient with this name already exists")
	}

	return core.client.Ingredient.Create().
		SetName(newIngredient.Name).
		SetCalories(newIngredient.Calories).
		SetFat(newIngredient.Fat).
		SetCarbohydrates(newIngredient.Carbohydrates).
		SetProtein(newIngredient.Protein).
		Save(c)
}

// Update an ingredient
func (core *IngredientCore) Update(name string, update Ingredient, c context.Context) (*ent.Ingredient, error) {
	previous, err := core.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(name)).
		First(c)

	if err != nil {
		return &ent.Ingredient{}, err
	}

	return core.client.Ingredient.UpdateOne(previous).
		SetName(update.Name).
		SetCalories(update.Calories).
		SetFat(update.Fat).
		SetCarbohydrates(update.Carbohydrates).
		SetProtein(update.Protein).
		Save(c)
}

// Get an ingredient by id
func (core *IngredientCore) Get(id string, c context.Context) (*ent.Ingredient, error) {
	return core.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(id)).
		First(c)
}

// GetAll ingredients within the limit/offset
func (core *IngredientCore) GetAll(limit int, offset int, c context.Context) (int, []*ent.Ingredient, error) {
	if limit == 0 || limit > 1000 {
		limit = 1000
	}

	count, err := core.client.
		Ingredient.
		Query().
		Count(c)

	if err != nil {
		return 0, make([]*ent.Ingredient, 0), err
	}

	if count == 0 || count < offset {
		return count, make([]*ent.Ingredient, 0), err
	}

	result, err := core.client.Ingredient.Query().
		Limit(limit).
		Offset(offset).
		All(c)

	return count, result, err
}

// Delete an ingredient
func (core *IngredientCore) Delete(name string, c context.Context) (*ent.Ingredient, error) {
	result, err := core.client.Ingredient.Query().
		Where(ingredient.NameEqualFold(name)).
		First(c)

	if err != nil {
		return &ent.Ingredient{}, err
	}

	err = core.client.Ingredient.DeleteOne(result).Exec(c)
	return result, err
}
