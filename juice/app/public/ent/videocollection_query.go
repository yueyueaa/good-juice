// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"juice/app/public/ent/predicate"
	"juice/app/public/ent/videocollection"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCollectionQuery is the builder for querying VideoCollection entities.
type VideoCollectionQuery struct {
	config
	ctx        *QueryContext
	order      []videocollection.OrderOption
	inters     []Interceptor
	predicates []predicate.VideoCollection
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoCollectionQuery builder.
func (vcq *VideoCollectionQuery) Where(ps ...predicate.VideoCollection) *VideoCollectionQuery {
	vcq.predicates = append(vcq.predicates, ps...)
	return vcq
}

// Limit the number of records to be returned by this query.
func (vcq *VideoCollectionQuery) Limit(limit int) *VideoCollectionQuery {
	vcq.ctx.Limit = &limit
	return vcq
}

// Offset to start from.
func (vcq *VideoCollectionQuery) Offset(offset int) *VideoCollectionQuery {
	vcq.ctx.Offset = &offset
	return vcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vcq *VideoCollectionQuery) Unique(unique bool) *VideoCollectionQuery {
	vcq.ctx.Unique = &unique
	return vcq
}

// Order specifies how the records should be ordered.
func (vcq *VideoCollectionQuery) Order(o ...videocollection.OrderOption) *VideoCollectionQuery {
	vcq.order = append(vcq.order, o...)
	return vcq
}

// First returns the first VideoCollection entity from the query.
// Returns a *NotFoundError when no VideoCollection was found.
func (vcq *VideoCollectionQuery) First(ctx context.Context) (*VideoCollection, error) {
	nodes, err := vcq.Limit(1).All(setContextOp(ctx, vcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videocollection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vcq *VideoCollectionQuery) FirstX(ctx context.Context) *VideoCollection {
	node, err := vcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VideoCollection ID from the query.
// Returns a *NotFoundError when no VideoCollection ID was found.
func (vcq *VideoCollectionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = vcq.Limit(1).IDs(setContextOp(ctx, vcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videocollection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vcq *VideoCollectionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := vcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VideoCollection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VideoCollection entity is found.
// Returns a *NotFoundError when no VideoCollection entities are found.
func (vcq *VideoCollectionQuery) Only(ctx context.Context) (*VideoCollection, error) {
	nodes, err := vcq.Limit(2).All(setContextOp(ctx, vcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videocollection.Label}
	default:
		return nil, &NotSingularError{videocollection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vcq *VideoCollectionQuery) OnlyX(ctx context.Context) *VideoCollection {
	node, err := vcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VideoCollection ID in the query.
// Returns a *NotSingularError when more than one VideoCollection ID is found.
// Returns a *NotFoundError when no entities are found.
func (vcq *VideoCollectionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = vcq.Limit(2).IDs(setContextOp(ctx, vcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videocollection.Label}
	default:
		err = &NotSingularError{videocollection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vcq *VideoCollectionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := vcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideoCollections.
func (vcq *VideoCollectionQuery) All(ctx context.Context) ([]*VideoCollection, error) {
	ctx = setContextOp(ctx, vcq.ctx, "All")
	if err := vcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VideoCollection, *VideoCollectionQuery]()
	return withInterceptors[[]*VideoCollection](ctx, vcq, qr, vcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vcq *VideoCollectionQuery) AllX(ctx context.Context) []*VideoCollection {
	nodes, err := vcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VideoCollection IDs.
func (vcq *VideoCollectionQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if vcq.ctx.Unique == nil && vcq.path != nil {
		vcq.Unique(true)
	}
	ctx = setContextOp(ctx, vcq.ctx, "IDs")
	if err = vcq.Select(videocollection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vcq *VideoCollectionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := vcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vcq *VideoCollectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vcq.ctx, "Count")
	if err := vcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vcq, querierCount[*VideoCollectionQuery](), vcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vcq *VideoCollectionQuery) CountX(ctx context.Context) int {
	count, err := vcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vcq *VideoCollectionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vcq.ctx, "Exist")
	switch _, err := vcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vcq *VideoCollectionQuery) ExistX(ctx context.Context) bool {
	exist, err := vcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoCollectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vcq *VideoCollectionQuery) Clone() *VideoCollectionQuery {
	if vcq == nil {
		return nil
	}
	return &VideoCollectionQuery{
		config:     vcq.config,
		ctx:        vcq.ctx.Clone(),
		order:      append([]videocollection.OrderOption{}, vcq.order...),
		inters:     append([]Interceptor{}, vcq.inters...),
		predicates: append([]predicate.VideoCollection{}, vcq.predicates...),
		// clone intermediate query.
		sql:  vcq.sql.Clone(),
		path: vcq.path,
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
//	client.VideoCollection.Query().
//		GroupBy(videocollection.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vcq *VideoCollectionQuery) GroupBy(field string, fields ...string) *VideoCollectionGroupBy {
	vcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoCollectionGroupBy{build: vcq}
	grbuild.flds = &vcq.ctx.Fields
	grbuild.label = videocollection.Label
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
//	client.VideoCollection.Query().
//		Select(videocollection.FieldUserID).
//		Scan(ctx, &v)
func (vcq *VideoCollectionQuery) Select(fields ...string) *VideoCollectionSelect {
	vcq.ctx.Fields = append(vcq.ctx.Fields, fields...)
	sbuild := &VideoCollectionSelect{VideoCollectionQuery: vcq}
	sbuild.label = videocollection.Label
	sbuild.flds, sbuild.scan = &vcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoCollectionSelect configured with the given aggregations.
func (vcq *VideoCollectionQuery) Aggregate(fns ...AggregateFunc) *VideoCollectionSelect {
	return vcq.Select().Aggregate(fns...)
}

func (vcq *VideoCollectionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vcq); err != nil {
				return err
			}
		}
	}
	for _, f := range vcq.ctx.Fields {
		if !videocollection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vcq.path != nil {
		prev, err := vcq.path(ctx)
		if err != nil {
			return err
		}
		vcq.sql = prev
	}
	return nil
}

func (vcq *VideoCollectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VideoCollection, error) {
	var (
		nodes = []*VideoCollection{}
		_spec = vcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VideoCollection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VideoCollection{config: vcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vcq *VideoCollectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vcq.querySpec()
	_spec.Node.Columns = vcq.ctx.Fields
	if len(vcq.ctx.Fields) > 0 {
		_spec.Unique = vcq.ctx.Unique != nil && *vcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vcq.driver, _spec)
}

func (vcq *VideoCollectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videocollection.Table, videocollection.Columns, sqlgraph.NewFieldSpec(videocollection.FieldID, field.TypeUint64))
	_spec.From = vcq.sql
	if unique := vcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vcq.path != nil {
		_spec.Unique = true
	}
	if fields := vcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videocollection.FieldID)
		for i := range fields {
			if fields[i] != videocollection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vcq *VideoCollectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vcq.driver.Dialect())
	t1 := builder.Table(videocollection.Table)
	columns := vcq.ctx.Fields
	if len(columns) == 0 {
		columns = videocollection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vcq.sql != nil {
		selector = vcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vcq.ctx.Unique != nil && *vcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vcq.predicates {
		p(selector)
	}
	for _, p := range vcq.order {
		p(selector)
	}
	if offset := vcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VideoCollectionGroupBy is the group-by builder for VideoCollection entities.
type VideoCollectionGroupBy struct {
	selector
	build *VideoCollectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vcgb *VideoCollectionGroupBy) Aggregate(fns ...AggregateFunc) *VideoCollectionGroupBy {
	vcgb.fns = append(vcgb.fns, fns...)
	return vcgb
}

// Scan applies the selector query and scans the result into the given value.
func (vcgb *VideoCollectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcgb.build.ctx, "GroupBy")
	if err := vcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCollectionQuery, *VideoCollectionGroupBy](ctx, vcgb.build, vcgb, vcgb.build.inters, v)
}

func (vcgb *VideoCollectionGroupBy) sqlScan(ctx context.Context, root *VideoCollectionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vcgb.fns))
	for _, fn := range vcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vcgb.flds)+len(vcgb.fns))
		for _, f := range *vcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoCollectionSelect is the builder for selecting fields of VideoCollection entities.
type VideoCollectionSelect struct {
	*VideoCollectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vcs *VideoCollectionSelect) Aggregate(fns ...AggregateFunc) *VideoCollectionSelect {
	vcs.fns = append(vcs.fns, fns...)
	return vcs
}

// Scan applies the selector query and scans the result into the given value.
func (vcs *VideoCollectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcs.ctx, "Select")
	if err := vcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCollectionQuery, *VideoCollectionSelect](ctx, vcs.VideoCollectionQuery, vcs, vcs.inters, v)
}

func (vcs *VideoCollectionSelect) sqlScan(ctx context.Context, root *VideoCollectionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vcs.fns))
	for _, fn := range vcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
