// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/ent/recipe"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecipeCreate is the builder for creating a Recipe entity.
type RecipeCreate struct {
	config
	mutation *RecipeMutation
	hooks    []Hook
}

// SetSlug sets the "slug" field.
func (rc *RecipeCreate) SetSlug(s string) *RecipeCreate {
	rc.mutation.SetSlug(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RecipeCreate) SetName(s string) *RecipeCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetIngredientslist sets the "ingredientslist" field.
func (rc *RecipeCreate) SetIngredientslist(s string) *RecipeCreate {
	rc.mutation.SetIngredientslist(s)
	return rc
}

// SetInstructions sets the "instructions" field.
func (rc *RecipeCreate) SetInstructions(s string) *RecipeCreate {
	rc.mutation.SetInstructions(s)
	return rc
}

// SetNutrition sets the "nutrition" field.
func (rc *RecipeCreate) SetNutrition(s string) *RecipeCreate {
	rc.mutation.SetNutrition(s)
	return rc
}

// SetUser sets the "user" field.
func (rc *RecipeCreate) SetUser(s string) *RecipeCreate {
	rc.mutation.SetUser(s)
	return rc
}

// SetServings sets the "servings" field.
func (rc *RecipeCreate) SetServings(i int) *RecipeCreate {
	rc.mutation.SetServings(i)
	return rc
}

// SetID sets the "id" field.
func (rc *RecipeCreate) SetID(s string) *RecipeCreate {
	rc.mutation.SetID(s)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableID(s *string) *RecipeCreate {
	if s != nil {
		rc.SetID(*s)
	}
	return rc
}

// AddIngredientIDs adds the "ingredients" edge to the Ingredient entity by IDs.
func (rc *RecipeCreate) AddIngredientIDs(ids ...int) *RecipeCreate {
	rc.mutation.AddIngredientIDs(ids...)
	return rc
}

// AddIngredients adds the "ingredients" edges to the Ingredient entity.
func (rc *RecipeCreate) AddIngredients(i ...*Ingredient) *RecipeCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return rc.AddIngredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (rc *RecipeCreate) Mutation() *RecipeMutation {
	return rc.mutation
}

// Save creates the Recipe in the database.
func (rc *RecipeCreate) Save(ctx context.Context) (*Recipe, error) {
	var (
		err  error
		node *Recipe
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecipeCreate) SaveX(ctx context.Context) *Recipe {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecipeCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecipeCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RecipeCreate) defaults() {
	if _, ok := rc.mutation.ID(); !ok {
		v := recipe.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecipeCreate) check() error {
	if _, ok := rc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Recipe.slug"`)}
	}
	if v, ok := rc.mutation.Slug(); ok {
		if err := recipe.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Recipe.slug": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Recipe.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := recipe.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Recipe.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Ingredientslist(); !ok {
		return &ValidationError{Name: "ingredientslist", err: errors.New(`ent: missing required field "Recipe.ingredientslist"`)}
	}
	if _, ok := rc.mutation.Instructions(); !ok {
		return &ValidationError{Name: "instructions", err: errors.New(`ent: missing required field "Recipe.instructions"`)}
	}
	if _, ok := rc.mutation.Nutrition(); !ok {
		return &ValidationError{Name: "nutrition", err: errors.New(`ent: missing required field "Recipe.nutrition"`)}
	}
	if _, ok := rc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "Recipe.user"`)}
	}
	if _, ok := rc.mutation.Servings(); !ok {
		return &ValidationError{Name: "servings", err: errors.New(`ent: missing required field "Recipe.servings"`)}
	}
	if v, ok := rc.mutation.Servings(); ok {
		if err := recipe.ServingsValidator(v); err != nil {
			return &ValidationError{Name: "servings", err: fmt.Errorf(`ent: validator failed for field "Recipe.servings": %w`, err)}
		}
	}
	return nil
}

func (rc *RecipeCreate) sqlSave(ctx context.Context) (*Recipe, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Recipe.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (rc *RecipeCreate) createSpec() (*Recipe, *sqlgraph.CreateSpec) {
	var (
		_node = &Recipe{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recipe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recipe.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldSlug,
		})
		_node.Slug = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Ingredientslist(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldIngredientslist,
		})
		_node.Ingredientslist = value
	}
	if value, ok := rc.mutation.Instructions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldInstructions,
		})
		_node.Instructions = value
	}
	if value, ok := rc.mutation.Nutrition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldNutrition,
		})
		_node.Nutrition = value
	}
	if value, ok := rc.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldUser,
		})
		_node.User = value
	}
	if value, ok := rc.mutation.Servings(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
		_node.Servings = value
	}
	if nodes := rc.mutation.IngredientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.IngredientsTable,
			Columns: recipe.IngredientsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecipeCreateBulk is the builder for creating many Recipe entities in bulk.
type RecipeCreateBulk struct {
	config
	builders []*RecipeCreate
}

// Save creates the Recipe entities in the database.
func (rcb *RecipeCreateBulk) Save(ctx context.Context) ([]*Recipe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recipe, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecipeCreateBulk) SaveX(ctx context.Context) []*Recipe {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecipeCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecipeCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
