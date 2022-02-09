// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"adomeit.xyz/recipe/ent/migrate"

	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/ent/recipe"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Ingredient is the client for interacting with the Ingredient builders.
	Ingredient *IngredientClient
	// Recipe is the client for interacting with the Recipe builders.
	Recipe *RecipeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Ingredient = NewIngredientClient(c.config)
	c.Recipe = NewRecipeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Ingredient: NewIngredientClient(cfg),
		Recipe:     NewRecipeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Ingredient: NewIngredientClient(cfg),
		Recipe:     NewRecipeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Ingredient.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Ingredient.Use(hooks...)
	c.Recipe.Use(hooks...)
}

// IngredientClient is a client for the Ingredient schema.
type IngredientClient struct {
	config
}

// NewIngredientClient returns a client for the Ingredient from the given config.
func NewIngredientClient(c config) *IngredientClient {
	return &IngredientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ingredient.Hooks(f(g(h())))`.
func (c *IngredientClient) Use(hooks ...Hook) {
	c.hooks.Ingredient = append(c.hooks.Ingredient, hooks...)
}

// Create returns a create builder for Ingredient.
func (c *IngredientClient) Create() *IngredientCreate {
	mutation := newIngredientMutation(c.config, OpCreate)
	return &IngredientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ingredient entities.
func (c *IngredientClient) CreateBulk(builders ...*IngredientCreate) *IngredientCreateBulk {
	return &IngredientCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ingredient.
func (c *IngredientClient) Update() *IngredientUpdate {
	mutation := newIngredientMutation(c.config, OpUpdate)
	return &IngredientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IngredientClient) UpdateOne(i *Ingredient) *IngredientUpdateOne {
	mutation := newIngredientMutation(c.config, OpUpdateOne, withIngredient(i))
	return &IngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IngredientClient) UpdateOneID(id int) *IngredientUpdateOne {
	mutation := newIngredientMutation(c.config, OpUpdateOne, withIngredientID(id))
	return &IngredientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ingredient.
func (c *IngredientClient) Delete() *IngredientDelete {
	mutation := newIngredientMutation(c.config, OpDelete)
	return &IngredientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *IngredientClient) DeleteOne(i *Ingredient) *IngredientDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *IngredientClient) DeleteOneID(id int) *IngredientDeleteOne {
	builder := c.Delete().Where(ingredient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IngredientDeleteOne{builder}
}

// Query returns a query builder for Ingredient.
func (c *IngredientClient) Query() *IngredientQuery {
	return &IngredientQuery{
		config: c.config,
	}
}

// Get returns a Ingredient entity by its id.
func (c *IngredientClient) Get(ctx context.Context, id int) (*Ingredient, error) {
	return c.Query().Where(ingredient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IngredientClient) GetX(ctx context.Context, id int) *Ingredient {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRecipe queries the recipe edge of a Ingredient.
func (c *IngredientClient) QueryRecipe(i *Ingredient) *RecipeQuery {
	query := &RecipeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ingredient.Table, ingredient.FieldID, id),
			sqlgraph.To(recipe.Table, recipe.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ingredient.RecipeTable, ingredient.RecipePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IngredientClient) Hooks() []Hook {
	return c.hooks.Ingredient
}

// RecipeClient is a client for the Recipe schema.
type RecipeClient struct {
	config
}

// NewRecipeClient returns a client for the Recipe from the given config.
func NewRecipeClient(c config) *RecipeClient {
	return &RecipeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `recipe.Hooks(f(g(h())))`.
func (c *RecipeClient) Use(hooks ...Hook) {
	c.hooks.Recipe = append(c.hooks.Recipe, hooks...)
}

// Create returns a create builder for Recipe.
func (c *RecipeClient) Create() *RecipeCreate {
	mutation := newRecipeMutation(c.config, OpCreate)
	return &RecipeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Recipe entities.
func (c *RecipeClient) CreateBulk(builders ...*RecipeCreate) *RecipeCreateBulk {
	return &RecipeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Recipe.
func (c *RecipeClient) Update() *RecipeUpdate {
	mutation := newRecipeMutation(c.config, OpUpdate)
	return &RecipeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RecipeClient) UpdateOne(r *Recipe) *RecipeUpdateOne {
	mutation := newRecipeMutation(c.config, OpUpdateOne, withRecipe(r))
	return &RecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RecipeClient) UpdateOneID(id string) *RecipeUpdateOne {
	mutation := newRecipeMutation(c.config, OpUpdateOne, withRecipeID(id))
	return &RecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Recipe.
func (c *RecipeClient) Delete() *RecipeDelete {
	mutation := newRecipeMutation(c.config, OpDelete)
	return &RecipeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RecipeClient) DeleteOne(r *Recipe) *RecipeDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RecipeClient) DeleteOneID(id string) *RecipeDeleteOne {
	builder := c.Delete().Where(recipe.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RecipeDeleteOne{builder}
}

// Query returns a query builder for Recipe.
func (c *RecipeClient) Query() *RecipeQuery {
	return &RecipeQuery{
		config: c.config,
	}
}

// Get returns a Recipe entity by its id.
func (c *RecipeClient) Get(ctx context.Context, id string) (*Recipe, error) {
	return c.Query().Where(recipe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RecipeClient) GetX(ctx context.Context, id string) *Recipe {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIngredients queries the ingredients edge of a Recipe.
func (c *RecipeClient) QueryIngredients(r *Recipe) *IngredientQuery {
	query := &IngredientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(recipe.Table, recipe.FieldID, id),
			sqlgraph.To(ingredient.Table, ingredient.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, recipe.IngredientsTable, recipe.IngredientsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RecipeClient) Hooks() []Hook {
	return c.hooks.Recipe
}
