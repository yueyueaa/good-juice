// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"juice/app/public/ent/predicate"
	"juice/app/public/ent/videolike"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoLikeQuery is the builder for querying VideoLike entities.
type VideoLikeQuery struct {
	config
	ctx        *QueryContext
	order      []videolike.OrderOption
	inters     []Interceptor
	predicates []predicate.VideoLike
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoLikeQuery builder.
func (vlq *VideoLikeQuery) Where(ps ...predicate.VideoLike) *VideoLikeQuery {
	vlq.predicates = append(vlq.predicates, ps...)
	return vlq
}

// Limit the number of records to be returned by this query.
func (vlq *VideoLikeQuery) Limit(limit int) *VideoLikeQuery {
	vlq.ctx.Limit = &limit
	return vlq
}

// Offset to start from.
func (vlq *VideoLikeQuery) Offset(offset int) *VideoLikeQuery {
	vlq.ctx.Offset = &offset
	return vlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vlq *VideoLikeQuery) Unique(unique bool) *VideoLikeQuery {
	vlq.ctx.Unique = &unique
	return vlq
}

// Order specifies how the records should be ordered.
func (vlq *VideoLikeQuery) Order(o ...videolike.OrderOption) *VideoLikeQuery {
	vlq.order = append(vlq.order, o...)
	return vlq
}

// First returns the first VideoLike entity from the query.
// Returns a *NotFoundError when no VideoLike was found.
func (vlq *VideoLikeQuery) First(ctx context.Context) (*VideoLike, error) {
	nodes, err := vlq.Limit(1).All(setContextOp(ctx, vlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videolike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vlq *VideoLikeQuery) FirstX(ctx context.Context) *VideoLike {
	node, err := vlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VideoLike ID from the query.
// Returns a *NotFoundError when no VideoLike ID was found.
func (vlq *VideoLikeQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = vlq.Limit(1).IDs(setContextOp(ctx, vlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videolike.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vlq *VideoLikeQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := vlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VideoLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VideoLike entity is found.
// Returns a *NotFoundError when no VideoLike entities are found.
func (vlq *VideoLikeQuery) Only(ctx context.Context) (*VideoLike, error) {
	nodes, err := vlq.Limit(2).All(setContextOp(ctx, vlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videolike.Label}
	default:
		return nil, &NotSingularError{videolike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vlq *VideoLikeQuery) OnlyX(ctx context.Context) *VideoLike {
	node, err := vlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VideoLike ID in the query.
// Returns a *NotSingularError when more than one VideoLike ID is found.
// Returns a *NotFoundError when no entities are found.
func (vlq *VideoLikeQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = vlq.Limit(2).IDs(setContextOp(ctx, vlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videolike.Label}
	default:
		err = &NotSingularError{videolike.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vlq *VideoLikeQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := vlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideoLikes.
func (vlq *VideoLikeQuery) All(ctx context.Context) ([]*VideoLike, error) {
	ctx = setContextOp(ctx, vlq.ctx, "All")
	if err := vlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VideoLike, *VideoLikeQuery]()
	return withInterceptors[[]*VideoLike](ctx, vlq, qr, vlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vlq *VideoLikeQuery) AllX(ctx context.Context) []*VideoLike {
	nodes, err := vlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VideoLike IDs.
func (vlq *VideoLikeQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if vlq.ctx.Unique == nil && vlq.path != nil {
		vlq.Unique(true)
	}
	ctx = setContextOp(ctx, vlq.ctx, "IDs")
	if err = vlq.Select(videolike.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vlq *VideoLikeQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := vlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vlq *VideoLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vlq.ctx, "Count")
	if err := vlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vlq, querierCount[*VideoLikeQuery](), vlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vlq *VideoLikeQuery) CountX(ctx context.Context) int {
	count, err := vlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vlq *VideoLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vlq.ctx, "Exist")
	switch _, err := vlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vlq *VideoLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := vlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vlq *VideoLikeQuery) Clone() *VideoLikeQuery {
	if vlq == nil {
		return nil
	}
	return &VideoLikeQuery{
		config:     vlq.config,
		ctx:        vlq.ctx.Clone(),
		order:      append([]videolike.OrderOption{}, vlq.order...),
		inters:     append([]Interceptor{}, vlq.inters...),
		predicates: append([]predicate.VideoLike{}, vlq.predicates...),
		// clone intermediate query.
		sql:  vlq.sql.Clone(),
		path: vlq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VideoLike.Query().
//		GroupBy(videolike.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vlq *VideoLikeQuery) GroupBy(field string, fields ...string) *VideoLikeGroupBy {
	vlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoLikeGroupBy{build: vlq}
	grbuild.flds = &vlq.ctx.Fields
	grbuild.label = videolike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.VideoLike.Query().
//		Select(videolike.FieldUserID).
//		Scan(ctx, &v)
func (vlq *VideoLikeQuery) Select(fields ...string) *VideoLikeSelect {
	vlq.ctx.Fields = append(vlq.ctx.Fields, fields...)
	sbuild := &VideoLikeSelect{VideoLikeQuery: vlq}
	sbuild.label = videolike.Label
	sbuild.flds, sbuild.scan = &vlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoLikeSelect configured with the given aggregations.
func (vlq *VideoLikeQuery) Aggregate(fns ...AggregateFunc) *VideoLikeSelect {
	return vlq.Select().Aggregate(fns...)
}

func (vlq *VideoLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vlq); err != nil {
				return err
			}
		}
	}
	for _, f := range vlq.ctx.Fields {
		if !videolike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vlq.path != nil {
		prev, err := vlq.path(ctx)
		if err != nil {
			return err
		}
		vlq.sql = prev
	}
	return nil
}

func (vlq *VideoLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VideoLike, error) {
	var (
		nodes = []*VideoLike{}
		_spec = vlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VideoLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VideoLike{config: vlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vlq *VideoLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vlq.querySpec()
	_spec.Node.Columns = vlq.ctx.Fields
	if len(vlq.ctx.Fields) > 0 {
		_spec.Unique = vlq.ctx.Unique != nil && *vlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vlq.driver, _spec)
}

func (vlq *VideoLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videolike.Table, videolike.Columns, sqlgraph.NewFieldSpec(videolike.FieldID, field.TypeUint64))
	_spec.From = vlq.sql
	if unique := vlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vlq.path != nil {
		_spec.Unique = true
	}
	if fields := vlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videolike.FieldID)
		for i := range fields {
			if fields[i] != videolike.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vlq *VideoLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vlq.driver.Dialect())
	t1 := builder.Table(videolike.Table)
	columns := vlq.ctx.Fields
	if len(columns) == 0 {
		columns = videolike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vlq.sql != nil {
		selector = vlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vlq.ctx.Unique != nil && *vlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vlq.predicates {
		p(selector)
	}
	for _, p := range vlq.order {
		p(selector)
	}
	if offset := vlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoLikeGroupBy is the group-by builder for VideoLike entities.
type VideoLikeGroupBy struct {
	selector
	build *VideoLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vlgb *VideoLikeGroupBy) Aggregate(fns ...AggregateFunc) *VideoLikeGroupBy {
	vlgb.fns = append(vlgb.fns, fns...)
	return vlgb
}

// Scan applies the selector query and scans the result into the given value.
func (vlgb *VideoLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vlgb.build.ctx, "GroupBy")
	if err := vlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoLikeQuery, *VideoLikeGroupBy](ctx, vlgb.build, vlgb, vlgb.build.inters, v)
}

func (vlgb *VideoLikeGroupBy) sqlScan(ctx context.Context, root *VideoLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vlgb.fns))
	for _, fn := range vlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vlgb.flds)+len(vlgb.fns))
		for _, f := range *vlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoLikeSelect is the builder for selecting fields of VideoLike entities.
type VideoLikeSelect struct {
	*VideoLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vls *VideoLikeSelect) Aggregate(fns ...AggregateFunc) *VideoLikeSelect {
	vls.fns = append(vls.fns, fns...)
	return vls
}

// Scan applies the selector query and scans the result into the given value.
func (vls *VideoLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vls.ctx, "Select")
	if err := vls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoLikeQuery, *VideoLikeSelect](ctx, vls.VideoLikeQuery, vls, vls.inters, v)
}

func (vls *VideoLikeSelect) sqlScan(ctx context.Context, root *VideoLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vls.fns))
	for _, fn := range vls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
