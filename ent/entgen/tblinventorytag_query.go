// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"
	"inventory/ent/entgen/predicate"
	"inventory/ent/entgen/tblinventory"
	"inventory/ent/entgen/tblinventorytag"
	"inventory/ent/entgen/tbltag"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblInventoryTagQuery is the builder for querying TblInventoryTag entities.
type TblInventoryTagQuery struct {
	config
	ctx           *QueryContext
	order         []tblinventorytag.OrderOption
	inters        []Interceptor
	predicates    []predicate.TblInventoryTag
	withInventory *TblInventoryQuery
	withTag       *TblTagQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TblInventoryTagQuery builder.
func (titq *TblInventoryTagQuery) Where(ps ...predicate.TblInventoryTag) *TblInventoryTagQuery {
	titq.predicates = append(titq.predicates, ps...)
	return titq
}

// Limit the number of records to be returned by this query.
func (titq *TblInventoryTagQuery) Limit(limit int) *TblInventoryTagQuery {
	titq.ctx.Limit = &limit
	return titq
}

// Offset to start from.
func (titq *TblInventoryTagQuery) Offset(offset int) *TblInventoryTagQuery {
	titq.ctx.Offset = &offset
	return titq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (titq *TblInventoryTagQuery) Unique(unique bool) *TblInventoryTagQuery {
	titq.ctx.Unique = &unique
	return titq
}

// Order specifies how the records should be ordered.
func (titq *TblInventoryTagQuery) Order(o ...tblinventorytag.OrderOption) *TblInventoryTagQuery {
	titq.order = append(titq.order, o...)
	return titq
}

// QueryInventory chains the current query on the "inventory" edge.
func (titq *TblInventoryTagQuery) QueryInventory() *TblInventoryQuery {
	query := (&TblInventoryClient{config: titq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := titq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := titq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblinventorytag.Table, tblinventorytag.FieldID, selector),
			sqlgraph.To(tblinventory.Table, tblinventory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tblinventorytag.InventoryTable, tblinventorytag.InventoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(titq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTag chains the current query on the "tag" edge.
func (titq *TblInventoryTagQuery) QueryTag() *TblTagQuery {
	query := (&TblTagClient{config: titq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := titq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := titq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblinventorytag.Table, tblinventorytag.FieldID, selector),
			sqlgraph.To(tbltag.Table, tbltag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tblinventorytag.TagTable, tblinventorytag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(titq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TblInventoryTag entity from the query.
// Returns a *NotFoundError when no TblInventoryTag was found.
func (titq *TblInventoryTagQuery) First(ctx context.Context) (*TblInventoryTag, error) {
	nodes, err := titq.Limit(1).All(setContextOp(ctx, titq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tblinventorytag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (titq *TblInventoryTagQuery) FirstX(ctx context.Context) *TblInventoryTag {
	node, err := titq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TblInventoryTag ID from the query.
// Returns a *NotFoundError when no TblInventoryTag ID was found.
func (titq *TblInventoryTagQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = titq.Limit(1).IDs(setContextOp(ctx, titq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tblinventorytag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (titq *TblInventoryTagQuery) FirstIDX(ctx context.Context) string {
	id, err := titq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TblInventoryTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TblInventoryTag entity is found.
// Returns a *NotFoundError when no TblInventoryTag entities are found.
func (titq *TblInventoryTagQuery) Only(ctx context.Context) (*TblInventoryTag, error) {
	nodes, err := titq.Limit(2).All(setContextOp(ctx, titq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tblinventorytag.Label}
	default:
		return nil, &NotSingularError{tblinventorytag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (titq *TblInventoryTagQuery) OnlyX(ctx context.Context) *TblInventoryTag {
	node, err := titq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TblInventoryTag ID in the query.
// Returns a *NotSingularError when more than one TblInventoryTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (titq *TblInventoryTagQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = titq.Limit(2).IDs(setContextOp(ctx, titq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tblinventorytag.Label}
	default:
		err = &NotSingularError{tblinventorytag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (titq *TblInventoryTagQuery) OnlyIDX(ctx context.Context) string {
	id, err := titq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TblInventoryTags.
func (titq *TblInventoryTagQuery) All(ctx context.Context) ([]*TblInventoryTag, error) {
	ctx = setContextOp(ctx, titq.ctx, ent.OpQueryAll)
	if err := titq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TblInventoryTag, *TblInventoryTagQuery]()
	return withInterceptors[[]*TblInventoryTag](ctx, titq, qr, titq.inters)
}

// AllX is like All, but panics if an error occurs.
func (titq *TblInventoryTagQuery) AllX(ctx context.Context) []*TblInventoryTag {
	nodes, err := titq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TblInventoryTag IDs.
func (titq *TblInventoryTagQuery) IDs(ctx context.Context) (ids []string, err error) {
	if titq.ctx.Unique == nil && titq.path != nil {
		titq.Unique(true)
	}
	ctx = setContextOp(ctx, titq.ctx, ent.OpQueryIDs)
	if err = titq.Select(tblinventorytag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (titq *TblInventoryTagQuery) IDsX(ctx context.Context) []string {
	ids, err := titq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (titq *TblInventoryTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, titq.ctx, ent.OpQueryCount)
	if err := titq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, titq, querierCount[*TblInventoryTagQuery](), titq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (titq *TblInventoryTagQuery) CountX(ctx context.Context) int {
	count, err := titq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (titq *TblInventoryTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, titq.ctx, ent.OpQueryExist)
	switch _, err := titq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (titq *TblInventoryTagQuery) ExistX(ctx context.Context) bool {
	exist, err := titq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TblInventoryTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (titq *TblInventoryTagQuery) Clone() *TblInventoryTagQuery {
	if titq == nil {
		return nil
	}
	return &TblInventoryTagQuery{
		config:        titq.config,
		ctx:           titq.ctx.Clone(),
		order:         append([]tblinventorytag.OrderOption{}, titq.order...),
		inters:        append([]Interceptor{}, titq.inters...),
		predicates:    append([]predicate.TblInventoryTag{}, titq.predicates...),
		withInventory: titq.withInventory.Clone(),
		withTag:       titq.withTag.Clone(),
		// clone intermediate query.
		sql:  titq.sql.Clone(),
		path: titq.path,
	}
}

// WithInventory tells the query-builder to eager-load the nodes that are connected to
// the "inventory" edge. The optional arguments are used to configure the query builder of the edge.
func (titq *TblInventoryTagQuery) WithInventory(opts ...func(*TblInventoryQuery)) *TblInventoryTagQuery {
	query := (&TblInventoryClient{config: titq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	titq.withInventory = query
	return titq
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (titq *TblInventoryTagQuery) WithTag(opts ...func(*TblTagQuery)) *TblInventoryTagQuery {
	query := (&TblTagClient{config: titq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	titq.withTag = query
	return titq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InventoryId string `json:"InventoryId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TblInventoryTag.Query().
//		GroupBy(tblinventorytag.FieldInventoryId).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
func (titq *TblInventoryTagQuery) GroupBy(field string, fields ...string) *TblInventoryTagGroupBy {
	titq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TblInventoryTagGroupBy{build: titq}
	grbuild.flds = &titq.ctx.Fields
	grbuild.label = tblinventorytag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InventoryId string `json:"InventoryId,omitempty"`
//	}
//
//	client.TblInventoryTag.Query().
//		Select(tblinventorytag.FieldInventoryId).
//		Scan(ctx, &v)
func (titq *TblInventoryTagQuery) Select(fields ...string) *TblInventoryTagSelect {
	titq.ctx.Fields = append(titq.ctx.Fields, fields...)
	sbuild := &TblInventoryTagSelect{TblInventoryTagQuery: titq}
	sbuild.label = tblinventorytag.Label
	sbuild.flds, sbuild.scan = &titq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TblInventoryTagSelect configured with the given aggregations.
func (titq *TblInventoryTagQuery) Aggregate(fns ...AggregateFunc) *TblInventoryTagSelect {
	return titq.Select().Aggregate(fns...)
}

func (titq *TblInventoryTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range titq.inters {
		if inter == nil {
			return fmt.Errorf("entgen: uninitialized interceptor (forgotten import entgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, titq); err != nil {
				return err
			}
		}
	}
	for _, f := range titq.ctx.Fields {
		if !tblinventorytag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if titq.path != nil {
		prev, err := titq.path(ctx)
		if err != nil {
			return err
		}
		titq.sql = prev
	}
	return nil
}

func (titq *TblInventoryTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TblInventoryTag, error) {
	var (
		nodes       = []*TblInventoryTag{}
		_spec       = titq.querySpec()
		loadedTypes = [2]bool{
			titq.withInventory != nil,
			titq.withTag != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TblInventoryTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TblInventoryTag{config: titq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(titq.modifiers) > 0 {
		_spec.Modifiers = titq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, titq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := titq.withInventory; query != nil {
		if err := titq.loadInventory(ctx, query, nodes, nil,
			func(n *TblInventoryTag, e *TblInventory) { n.Edges.Inventory = e }); err != nil {
			return nil, err
		}
	}
	if query := titq.withTag; query != nil {
		if err := titq.loadTag(ctx, query, nodes, nil,
			func(n *TblInventoryTag, e *TblTag) { n.Edges.Tag = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (titq *TblInventoryTagQuery) loadInventory(ctx context.Context, query *TblInventoryQuery, nodes []*TblInventoryTag, init func(*TblInventoryTag), assign func(*TblInventoryTag, *TblInventory)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TblInventoryTag)
	for i := range nodes {
		fk := nodes[i].InventoryId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tblinventory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "InventoryId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (titq *TblInventoryTagQuery) loadTag(ctx context.Context, query *TblTagQuery, nodes []*TblInventoryTag, init func(*TblInventoryTag), assign func(*TblInventoryTag, *TblTag)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TblInventoryTag)
	for i := range nodes {
		fk := nodes[i].TagId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tbltag.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "TagId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (titq *TblInventoryTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := titq.querySpec()
	if len(titq.modifiers) > 0 {
		_spec.Modifiers = titq.modifiers
	}
	_spec.Node.Columns = titq.ctx.Fields
	if len(titq.ctx.Fields) > 0 {
		_spec.Unique = titq.ctx.Unique != nil && *titq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, titq.driver, _spec)
}

func (titq *TblInventoryTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tblinventorytag.Table, tblinventorytag.Columns, sqlgraph.NewFieldSpec(tblinventorytag.FieldID, field.TypeString))
	_spec.From = titq.sql
	if unique := titq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if titq.path != nil {
		_spec.Unique = true
	}
	if fields := titq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblinventorytag.FieldID)
		for i := range fields {
			if fields[i] != tblinventorytag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if titq.withInventory != nil {
			_spec.Node.AddColumnOnce(tblinventorytag.FieldInventoryId)
		}
		if titq.withTag != nil {
			_spec.Node.AddColumnOnce(tblinventorytag.FieldTagId)
		}
	}
	if ps := titq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := titq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := titq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := titq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (titq *TblInventoryTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(titq.driver.Dialect())
	t1 := builder.Table(tblinventorytag.Table)
	columns := titq.ctx.Fields
	if len(columns) == 0 {
		columns = tblinventorytag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if titq.sql != nil {
		selector = titq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if titq.ctx.Unique != nil && *titq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range titq.modifiers {
		m(selector)
	}
	for _, p := range titq.predicates {
		p(selector)
	}
	for _, p := range titq.order {
		p(selector)
	}
	if offset := titq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := titq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (titq *TblInventoryTagQuery) ForUpdate(opts ...sql.LockOption) *TblInventoryTagQuery {
	if titq.driver.Dialect() == dialect.Postgres {
		titq.Unique(false)
	}
	titq.modifiers = append(titq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return titq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (titq *TblInventoryTagQuery) ForShare(opts ...sql.LockOption) *TblInventoryTagQuery {
	if titq.driver.Dialect() == dialect.Postgres {
		titq.Unique(false)
	}
	titq.modifiers = append(titq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return titq
}

// TblInventoryTagGroupBy is the group-by builder for TblInventoryTag entities.
type TblInventoryTagGroupBy struct {
	selector
	build *TblInventoryTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (titgb *TblInventoryTagGroupBy) Aggregate(fns ...AggregateFunc) *TblInventoryTagGroupBy {
	titgb.fns = append(titgb.fns, fns...)
	return titgb
}

// Scan applies the selector query and scans the result into the given value.
func (titgb *TblInventoryTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, titgb.build.ctx, ent.OpQueryGroupBy)
	if err := titgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblInventoryTagQuery, *TblInventoryTagGroupBy](ctx, titgb.build, titgb, titgb.build.inters, v)
}

func (titgb *TblInventoryTagGroupBy) sqlScan(ctx context.Context, root *TblInventoryTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(titgb.fns))
	for _, fn := range titgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*titgb.flds)+len(titgb.fns))
		for _, f := range *titgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*titgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := titgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TblInventoryTagSelect is the builder for selecting fields of TblInventoryTag entities.
type TblInventoryTagSelect struct {
	*TblInventoryTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tits *TblInventoryTagSelect) Aggregate(fns ...AggregateFunc) *TblInventoryTagSelect {
	tits.fns = append(tits.fns, fns...)
	return tits
}

// Scan applies the selector query and scans the result into the given value.
func (tits *TblInventoryTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tits.ctx, ent.OpQuerySelect)
	if err := tits.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblInventoryTagQuery, *TblInventoryTagSelect](ctx, tits.TblInventoryTagQuery, tits, tits.inters, v)
}

func (tits *TblInventoryTagSelect) sqlScan(ctx context.Context, root *TblInventoryTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tits.fns))
	for _, fn := range tits.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tits.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tits.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
