// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"juice/internal/data/ent/predicate"
	"juice/internal/data/ent/videocomment"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCommentQuery is the builder for querying VideoComment entities.
type VideoCommentQuery struct {
	config
	ctx        *QueryContext
	order      []videocomment.OrderOption
	inters     []Interceptor
	predicates []predicate.VideoComment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoCommentQuery builder.
func (vcq *VideoCommentQuery) Where(ps ...predicate.VideoComment) *VideoCommentQuery {
	vcq.predicates = append(vcq.predicates, ps...)
	return vcq
}

// Limit the number of records to be returned by this query.
func (vcq *VideoCommentQuery) Limit(limit int) *VideoCommentQuery {
	vcq.ctx.Limit = &limit
	return vcq
}

// Offset to start from.
func (vcq *VideoCommentQuery) Offset(offset int) *VideoCommentQuery {
	vcq.ctx.Offset = &offset
	return vcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vcq *VideoCommentQuery) Unique(unique bool) *VideoCommentQuery {
	vcq.ctx.Unique = &unique
	return vcq
}

// Order specifies how the records should be ordered.
func (vcq *VideoCommentQuery) Order(o ...videocomment.OrderOption) *VideoCommentQuery {
	vcq.order = append(vcq.order, o...)
	return vcq
}

// First returns the first VideoComment entity from the query.
// Returns a *NotFoundError when no VideoComment was found.
func (vcq *VideoCommentQuery) First(ctx context.Context) (*VideoComment, error) {
	nodes, err := vcq.Limit(1).All(setContextOp(ctx, vcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videocomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vcq *VideoCommentQuery) FirstX(ctx context.Context) *VideoComment {
	node, err := vcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VideoComment ID from the query.
// Returns a *NotFoundError when no VideoComment ID was found.
func (vcq *VideoCommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vcq.Limit(1).IDs(setContextOp(ctx, vcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videocomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vcq *VideoCommentQuery) FirstIDX(ctx context.Context) int {
	id, err := vcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VideoComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VideoComment entity is found.
// Returns a *NotFoundError when no VideoComment entities are found.
func (vcq *VideoCommentQuery) Only(ctx context.Context) (*VideoComment, error) {
	nodes, err := vcq.Limit(2).All(setContextOp(ctx, vcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videocomment.Label}
	default:
		return nil, &NotSingularError{videocomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vcq *VideoCommentQuery) OnlyX(ctx context.Context) *VideoComment {
	node, err := vcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VideoComment ID in the query.
// Returns a *NotSingularError when more than one VideoComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (vcq *VideoCommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vcq.Limit(2).IDs(setContextOp(ctx, vcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videocomment.Label}
	default:
		err = &NotSingularError{videocomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vcq *VideoCommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := vcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideoComments.
func (vcq *VideoCommentQuery) All(ctx context.Context) ([]*VideoComment, error) {
	ctx = setContextOp(ctx, vcq.ctx, "All")
	if err := vcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VideoComment, *VideoCommentQuery]()
	return withInterceptors[[]*VideoComment](ctx, vcq, qr, vcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vcq *VideoCommentQuery) AllX(ctx context.Context) []*VideoComment {
	nodes, err := vcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VideoComment IDs.
func (vcq *VideoCommentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vcq.ctx.Unique == nil && vcq.path != nil {
		vcq.Unique(true)
	}
	ctx = setContextOp(ctx, vcq.ctx, "IDs")
	if err = vcq.Select(videocomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vcq *VideoCommentQuery) IDsX(ctx context.Context) []int {
	ids, err := vcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vcq *VideoCommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vcq.ctx, "Count")
	if err := vcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vcq, querierCount[*VideoCommentQuery](), vcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vcq *VideoCommentQuery) CountX(ctx context.Context) int {
	count, err := vcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vcq *VideoCommentQuery) Exist(ctx context.Context) (bool, error) {
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
func (vcq *VideoCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := vcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vcq *VideoCommentQuery) Clone() *VideoCommentQuery {
	if vcq == nil {
		return nil
	}
	return &VideoCommentQuery{
		config:     vcq.config,
		ctx:        vcq.ctx.Clone(),
		order:      append([]videocomment.OrderOption{}, vcq.order...),
		inters:     append([]Interceptor{}, vcq.inters...),
		predicates: append([]predicate.VideoComment{}, vcq.predicates...),
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
//		CommentID int `json:"comment_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VideoComment.Query().
//		GroupBy(videocomment.FieldCommentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vcq *VideoCommentQuery) GroupBy(field string, fields ...string) *VideoCommentGroupBy {
	vcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoCommentGroupBy{build: vcq}
	grbuild.flds = &vcq.ctx.Fields
	grbuild.label = videocomment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CommentID int `json:"comment_id,omitempty"`
//	}
//
//	client.VideoComment.Query().
//		Select(videocomment.FieldCommentID).
//		Scan(ctx, &v)
func (vcq *VideoCommentQuery) Select(fields ...string) *VideoCommentSelect {
	vcq.ctx.Fields = append(vcq.ctx.Fields, fields...)
	sbuild := &VideoCommentSelect{VideoCommentQuery: vcq}
	sbuild.label = videocomment.Label
	sbuild.flds, sbuild.scan = &vcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoCommentSelect configured with the given aggregations.
func (vcq *VideoCommentQuery) Aggregate(fns ...AggregateFunc) *VideoCommentSelect {
	return vcq.Select().Aggregate(fns...)
}

func (vcq *VideoCommentQuery) prepareQuery(ctx context.Context) error {
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
		if !videocomment.ValidColumn(f) {
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

func (vcq *VideoCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VideoComment, error) {
	var (
		nodes = []*VideoComment{}
		_spec = vcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VideoComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VideoComment{config: vcq.config}
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

func (vcq *VideoCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vcq.querySpec()
	_spec.Node.Columns = vcq.ctx.Fields
	if len(vcq.ctx.Fields) > 0 {
		_spec.Unique = vcq.ctx.Unique != nil && *vcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vcq.driver, _spec)
}

func (vcq *VideoCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videocomment.Table, videocomment.Columns, sqlgraph.NewFieldSpec(videocomment.FieldID, field.TypeInt))
	_spec.From = vcq.sql
	if unique := vcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vcq.path != nil {
		_spec.Unique = true
	}
	if fields := vcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videocomment.FieldID)
		for i := range fields {
			if fields[i] != videocomment.FieldID {
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

func (vcq *VideoCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vcq.driver.Dialect())
	t1 := builder.Table(videocomment.Table)
	columns := vcq.ctx.Fields
	if len(columns) == 0 {
		columns = videocomment.Columns
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

// VideoCommentGroupBy is the group-by builder for VideoComment entities.
type VideoCommentGroupBy struct {
	selector
	build *VideoCommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vcgb *VideoCommentGroupBy) Aggregate(fns ...AggregateFunc) *VideoCommentGroupBy {
	vcgb.fns = append(vcgb.fns, fns...)
	return vcgb
}

// Scan applies the selector query and scans the result into the given value.
func (vcgb *VideoCommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcgb.build.ctx, "GroupBy")
	if err := vcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCommentQuery, *VideoCommentGroupBy](ctx, vcgb.build, vcgb, vcgb.build.inters, v)
}

func (vcgb *VideoCommentGroupBy) sqlScan(ctx context.Context, root *VideoCommentQuery, v any) error {
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

// VideoCommentSelect is the builder for selecting fields of VideoComment entities.
type VideoCommentSelect struct {
	*VideoCommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vcs *VideoCommentSelect) Aggregate(fns ...AggregateFunc) *VideoCommentSelect {
	vcs.fns = append(vcs.fns, fns...)
	return vcs
}

// Scan applies the selector query and scans the result into the given value.
func (vcs *VideoCommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcs.ctx, "Select")
	if err := vcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCommentQuery, *VideoCommentSelect](ctx, vcs.VideoCommentQuery, vcs, vcs.inters, v)
}

func (vcs *VideoCommentSelect) sqlScan(ctx context.Context, root *VideoCommentQuery, v any) error {
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
