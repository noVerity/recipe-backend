// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"adomeit.xyz/recipe/ent/ingredient"
	"adomeit.xyz/recipe/ent/predicate"
	"adomeit.xyz/recipe/ent/recipe"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IngredientQuery is the builder for querying Ingredient entities.
type IngredientQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Ingredient
	// eager-loading edges.
	withRecipe *RecipeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IngredientQuery builder.
func (iq *IngredientQuery) Where(ps ...predicate.Ingredient) *IngredientQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *IngredientQuery) Limit(limit int) *IngredientQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *IngredientQuery) Offset(offset int) *IngredientQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IngredientQuery) Unique(unique bool) *IngredientQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *IngredientQuery) Order(o ...OrderFunc) *IngredientQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryRecipe chains the current query on the "recipe" edge.
func (iq *IngredientQuery) QueryRecipe() *RecipeQuery {
	query := &RecipeQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ingredient.Table, ingredient.FieldID, selector),
			sqlgraph.To(recipe.Table, recipe.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ingredient.RecipeTable, ingredient.RecipePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Ingredient entity from the query.
// Returns a *NotFoundError when no Ingredient was found.
func (iq *IngredientQuery) First(ctx context.Context) (*Ingredient, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ingredient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IngredientQuery) FirstX(ctx context.Context) *Ingredient {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ingredient ID from the query.
// Returns a *NotFoundError when no Ingredient ID was found.
func (iq *IngredientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ingredient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IngredientQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ingredient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Ingredient entity is not found.
// Returns a *NotFoundError when no Ingredient entities are found.
func (iq *IngredientQuery) Only(ctx context.Context) (*Ingredient, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ingredient.Label}
	default:
		return nil, &NotSingularError{ingredient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IngredientQuery) OnlyX(ctx context.Context) *Ingredient {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ingredient ID in the query.
// Returns a *NotSingularError when exactly one Ingredient ID is not found.
// Returns a *NotFoundError when no entities are found.
func (iq *IngredientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = &NotSingularError{ingredient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IngredientQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ingredients.
func (iq *IngredientQuery) All(ctx context.Context) ([]*Ingredient, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *IngredientQuery) AllX(ctx context.Context) []*Ingredient {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ingredient IDs.
func (iq *IngredientQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(ingredient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IngredientQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IngredientQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IngredientQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IngredientQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IngredientQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IngredientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IngredientQuery) Clone() *IngredientQuery {
	if iq == nil {
		return nil
	}
	return &IngredientQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		predicates: append([]predicate.Ingredient{}, iq.predicates...),
		withRecipe: iq.withRecipe.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithRecipe tells the query-builder to eager-load the nodes that are connected to
// the "recipe" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IngredientQuery) WithRecipe(opts ...func(*RecipeQuery)) *IngredientQuery {
	query := &RecipeQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withRecipe = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ingredient.Query().
//		GroupBy(ingredient.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *IngredientQuery) GroupBy(field string, fields ...string) *IngredientGroupBy {
	group := &IngredientGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Ingredient.Query().
//		Select(ingredient.FieldName).
//		Scan(ctx, &v)
//
func (iq *IngredientQuery) Select(fields ...string) *IngredientSelect {
	iq.fields = append(iq.fields, fields...)
	return &IngredientSelect{IngredientQuery: iq}
}

func (iq *IngredientQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !ingredient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IngredientQuery) sqlAll(ctx context.Context) ([]*Ingredient, error) {
	var (
		nodes       = []*Ingredient{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withRecipe != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Ingredient{config: iq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withRecipe; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Ingredient, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Recipe = []*Recipe{}
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Ingredient)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   ingredient.RecipeTable,
				Columns: ingredient.RecipePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(ingredient.RecipePrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullString)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullString)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := ein.String
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, iq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "recipe": %w`, err)
		}
		query.Where(recipe.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "recipe" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Recipe = append(nodes[i].Edges.Recipe, n)
			}
		}
	}

	return nodes, nil
}

func (iq *IngredientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IngredientQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *IngredientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ingredient.Table,
			Columns: ingredient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ingredient.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ingredient.FieldID)
		for i := range fields {
			if fields[i] != ingredient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IngredientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(ingredient.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = ingredient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IngredientGroupBy is the group-by builder for Ingredient entities.
type IngredientGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IngredientGroupBy) Aggregate(fns ...AggregateFunc) *IngredientGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *IngredientGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *IngredientGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IngredientGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *IngredientGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = igb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (igb *IngredientGroupBy) StringX(ctx context.Context) string {
	v, err := igb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IngredientGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *IngredientGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = igb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (igb *IngredientGroupBy) IntX(ctx context.Context) int {
	v, err := igb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IngredientGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *IngredientGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = igb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (igb *IngredientGroupBy) Float64X(ctx context.Context) float64 {
	v, err := igb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: IngredientGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *IngredientGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *IngredientGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = igb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (igb *IngredientGroupBy) BoolX(ctx context.Context) bool {
	v, err := igb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *IngredientGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !ingredient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *IngredientGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// IngredientSelect is the builder for selecting fields of Ingredient entities.
type IngredientSelect struct {
	*IngredientQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *IngredientSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.IngredientQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *IngredientSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IngredientSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *IngredientSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = is.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (is *IngredientSelect) StringX(ctx context.Context) string {
	v, err := is.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IngredientSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *IngredientSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = is.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (is *IngredientSelect) IntX(ctx context.Context) int {
	v, err := is.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IngredientSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *IngredientSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = is.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (is *IngredientSelect) Float64X(ctx context.Context) float64 {
	v, err := is.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: IngredientSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *IngredientSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (is *IngredientSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = is.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = fmt.Errorf("ent: IngredientSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (is *IngredientSelect) BoolX(ctx context.Context) bool {
	v, err := is.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *IngredientSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
