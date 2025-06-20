// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"inventory/ent/entgen/predicate"
	"inventory/ent/entgen/tblinventory"
	"inventory/ent/entgen/tblinventorytag"
	"inventory/ent/entgen/tbltag"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblInventoryTagUpdate is the builder for updating TblInventoryTag entities.
type TblInventoryTagUpdate struct {
	config
	hooks    []Hook
	mutation *TblInventoryTagMutation
}

// Where appends a list predicates to the TblInventoryTagUpdate builder.
func (titu *TblInventoryTagUpdate) Where(ps ...predicate.TblInventoryTag) *TblInventoryTagUpdate {
	titu.mutation.Where(ps...)
	return titu
}

// SetInventoryId sets the "InventoryId" field.
func (titu *TblInventoryTagUpdate) SetInventoryId(s string) *TblInventoryTagUpdate {
	titu.mutation.SetInventoryId(s)
	return titu
}

// SetNillableInventoryId sets the "InventoryId" field if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableInventoryId(s *string) *TblInventoryTagUpdate {
	if s != nil {
		titu.SetInventoryId(*s)
	}
	return titu
}

// ClearInventoryId clears the value of the "InventoryId" field.
func (titu *TblInventoryTagUpdate) ClearInventoryId() *TblInventoryTagUpdate {
	titu.mutation.ClearInventoryId()
	return titu
}

// SetTagId sets the "TagId" field.
func (titu *TblInventoryTagUpdate) SetTagId(s string) *TblInventoryTagUpdate {
	titu.mutation.SetTagId(s)
	return titu
}

// SetNillableTagId sets the "TagId" field if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableTagId(s *string) *TblInventoryTagUpdate {
	if s != nil {
		titu.SetTagId(*s)
	}
	return titu
}

// ClearTagId clears the value of the "TagId" field.
func (titu *TblInventoryTagUpdate) ClearTagId() *TblInventoryTagUpdate {
	titu.mutation.ClearTagId()
	return titu
}

// SetCreatedAt sets the "Created_at" field.
func (titu *TblInventoryTagUpdate) SetCreatedAt(t time.Time) *TblInventoryTagUpdate {
	titu.mutation.SetCreatedAt(t)
	return titu
}

// SetNillableCreatedAt sets the "Created_at" field if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableCreatedAt(t *time.Time) *TblInventoryTagUpdate {
	if t != nil {
		titu.SetCreatedAt(*t)
	}
	return titu
}

// SetUpdatedAt sets the "Updated_at" field.
func (titu *TblInventoryTagUpdate) SetUpdatedAt(t time.Time) *TblInventoryTagUpdate {
	titu.mutation.SetUpdatedAt(t)
	return titu
}

// SetNillableUpdatedAt sets the "Updated_at" field if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableUpdatedAt(t *time.Time) *TblInventoryTagUpdate {
	if t != nil {
		titu.SetUpdatedAt(*t)
	}
	return titu
}

// SetDeletedAt sets the "Deleted_at" field.
func (titu *TblInventoryTagUpdate) SetDeletedAt(t time.Time) *TblInventoryTagUpdate {
	titu.mutation.SetDeletedAt(t)
	return titu
}

// SetNillableDeletedAt sets the "Deleted_at" field if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableDeletedAt(t *time.Time) *TblInventoryTagUpdate {
	if t != nil {
		titu.SetDeletedAt(*t)
	}
	return titu
}

// ClearDeletedAt clears the value of the "Deleted_at" field.
func (titu *TblInventoryTagUpdate) ClearDeletedAt() *TblInventoryTagUpdate {
	titu.mutation.ClearDeletedAt()
	return titu
}

// SetInventoryID sets the "inventory" edge to the TblInventory entity by ID.
func (titu *TblInventoryTagUpdate) SetInventoryID(id string) *TblInventoryTagUpdate {
	titu.mutation.SetInventoryID(id)
	return titu
}

// SetNillableInventoryID sets the "inventory" edge to the TblInventory entity by ID if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableInventoryID(id *string) *TblInventoryTagUpdate {
	if id != nil {
		titu = titu.SetInventoryID(*id)
	}
	return titu
}

// SetInventory sets the "inventory" edge to the TblInventory entity.
func (titu *TblInventoryTagUpdate) SetInventory(t *TblInventory) *TblInventoryTagUpdate {
	return titu.SetInventoryID(t.ID)
}

// SetTagID sets the "tag" edge to the TblTag entity by ID.
func (titu *TblInventoryTagUpdate) SetTagID(id string) *TblInventoryTagUpdate {
	titu.mutation.SetTagID(id)
	return titu
}

// SetNillableTagID sets the "tag" edge to the TblTag entity by ID if the given value is not nil.
func (titu *TblInventoryTagUpdate) SetNillableTagID(id *string) *TblInventoryTagUpdate {
	if id != nil {
		titu = titu.SetTagID(*id)
	}
	return titu
}

// SetTag sets the "tag" edge to the TblTag entity.
func (titu *TblInventoryTagUpdate) SetTag(t *TblTag) *TblInventoryTagUpdate {
	return titu.SetTagID(t.ID)
}

// Mutation returns the TblInventoryTagMutation object of the builder.
func (titu *TblInventoryTagUpdate) Mutation() *TblInventoryTagMutation {
	return titu.mutation
}

// ClearInventory clears the "inventory" edge to the TblInventory entity.
func (titu *TblInventoryTagUpdate) ClearInventory() *TblInventoryTagUpdate {
	titu.mutation.ClearInventory()
	return titu
}

// ClearTag clears the "tag" edge to the TblTag entity.
func (titu *TblInventoryTagUpdate) ClearTag() *TblInventoryTagUpdate {
	titu.mutation.ClearTag()
	return titu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (titu *TblInventoryTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, titu.sqlSave, titu.mutation, titu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (titu *TblInventoryTagUpdate) SaveX(ctx context.Context) int {
	affected, err := titu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (titu *TblInventoryTagUpdate) Exec(ctx context.Context) error {
	_, err := titu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (titu *TblInventoryTagUpdate) ExecX(ctx context.Context) {
	if err := titu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (titu *TblInventoryTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tblinventorytag.Table, tblinventorytag.Columns, sqlgraph.NewFieldSpec(tblinventorytag.FieldID, field.TypeString))
	if ps := titu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := titu.mutation.CreatedAt(); ok {
		_spec.SetField(tblinventorytag.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := titu.mutation.UpdatedAt(); ok {
		_spec.SetField(tblinventorytag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := titu.mutation.DeletedAt(); ok {
		_spec.SetField(tblinventorytag.FieldDeletedAt, field.TypeTime, value)
	}
	if titu.mutation.DeletedAtCleared() {
		_spec.ClearField(tblinventorytag.FieldDeletedAt, field.TypeTime)
	}
	if titu.mutation.InventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.InventoryTable,
			Columns: []string{tblinventorytag.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := titu.mutation.InventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.InventoryTable,
			Columns: []string{tblinventorytag.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if titu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.TagTable,
			Columns: []string{tblinventorytag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbltag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := titu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.TagTable,
			Columns: []string{tblinventorytag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbltag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, titu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblinventorytag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	titu.mutation.done = true
	return n, nil
}

// TblInventoryTagUpdateOne is the builder for updating a single TblInventoryTag entity.
type TblInventoryTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TblInventoryTagMutation
}

// SetInventoryId sets the "InventoryId" field.
func (tituo *TblInventoryTagUpdateOne) SetInventoryId(s string) *TblInventoryTagUpdateOne {
	tituo.mutation.SetInventoryId(s)
	return tituo
}

// SetNillableInventoryId sets the "InventoryId" field if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableInventoryId(s *string) *TblInventoryTagUpdateOne {
	if s != nil {
		tituo.SetInventoryId(*s)
	}
	return tituo
}

// ClearInventoryId clears the value of the "InventoryId" field.
func (tituo *TblInventoryTagUpdateOne) ClearInventoryId() *TblInventoryTagUpdateOne {
	tituo.mutation.ClearInventoryId()
	return tituo
}

// SetTagId sets the "TagId" field.
func (tituo *TblInventoryTagUpdateOne) SetTagId(s string) *TblInventoryTagUpdateOne {
	tituo.mutation.SetTagId(s)
	return tituo
}

// SetNillableTagId sets the "TagId" field if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableTagId(s *string) *TblInventoryTagUpdateOne {
	if s != nil {
		tituo.SetTagId(*s)
	}
	return tituo
}

// ClearTagId clears the value of the "TagId" field.
func (tituo *TblInventoryTagUpdateOne) ClearTagId() *TblInventoryTagUpdateOne {
	tituo.mutation.ClearTagId()
	return tituo
}

// SetCreatedAt sets the "Created_at" field.
func (tituo *TblInventoryTagUpdateOne) SetCreatedAt(t time.Time) *TblInventoryTagUpdateOne {
	tituo.mutation.SetCreatedAt(t)
	return tituo
}

// SetNillableCreatedAt sets the "Created_at" field if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableCreatedAt(t *time.Time) *TblInventoryTagUpdateOne {
	if t != nil {
		tituo.SetCreatedAt(*t)
	}
	return tituo
}

// SetUpdatedAt sets the "Updated_at" field.
func (tituo *TblInventoryTagUpdateOne) SetUpdatedAt(t time.Time) *TblInventoryTagUpdateOne {
	tituo.mutation.SetUpdatedAt(t)
	return tituo
}

// SetNillableUpdatedAt sets the "Updated_at" field if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableUpdatedAt(t *time.Time) *TblInventoryTagUpdateOne {
	if t != nil {
		tituo.SetUpdatedAt(*t)
	}
	return tituo
}

// SetDeletedAt sets the "Deleted_at" field.
func (tituo *TblInventoryTagUpdateOne) SetDeletedAt(t time.Time) *TblInventoryTagUpdateOne {
	tituo.mutation.SetDeletedAt(t)
	return tituo
}

// SetNillableDeletedAt sets the "Deleted_at" field if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableDeletedAt(t *time.Time) *TblInventoryTagUpdateOne {
	if t != nil {
		tituo.SetDeletedAt(*t)
	}
	return tituo
}

// ClearDeletedAt clears the value of the "Deleted_at" field.
func (tituo *TblInventoryTagUpdateOne) ClearDeletedAt() *TblInventoryTagUpdateOne {
	tituo.mutation.ClearDeletedAt()
	return tituo
}

// SetInventoryID sets the "inventory" edge to the TblInventory entity by ID.
func (tituo *TblInventoryTagUpdateOne) SetInventoryID(id string) *TblInventoryTagUpdateOne {
	tituo.mutation.SetInventoryID(id)
	return tituo
}

// SetNillableInventoryID sets the "inventory" edge to the TblInventory entity by ID if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableInventoryID(id *string) *TblInventoryTagUpdateOne {
	if id != nil {
		tituo = tituo.SetInventoryID(*id)
	}
	return tituo
}

// SetInventory sets the "inventory" edge to the TblInventory entity.
func (tituo *TblInventoryTagUpdateOne) SetInventory(t *TblInventory) *TblInventoryTagUpdateOne {
	return tituo.SetInventoryID(t.ID)
}

// SetTagID sets the "tag" edge to the TblTag entity by ID.
func (tituo *TblInventoryTagUpdateOne) SetTagID(id string) *TblInventoryTagUpdateOne {
	tituo.mutation.SetTagID(id)
	return tituo
}

// SetNillableTagID sets the "tag" edge to the TblTag entity by ID if the given value is not nil.
func (tituo *TblInventoryTagUpdateOne) SetNillableTagID(id *string) *TblInventoryTagUpdateOne {
	if id != nil {
		tituo = tituo.SetTagID(*id)
	}
	return tituo
}

// SetTag sets the "tag" edge to the TblTag entity.
func (tituo *TblInventoryTagUpdateOne) SetTag(t *TblTag) *TblInventoryTagUpdateOne {
	return tituo.SetTagID(t.ID)
}

// Mutation returns the TblInventoryTagMutation object of the builder.
func (tituo *TblInventoryTagUpdateOne) Mutation() *TblInventoryTagMutation {
	return tituo.mutation
}

// ClearInventory clears the "inventory" edge to the TblInventory entity.
func (tituo *TblInventoryTagUpdateOne) ClearInventory() *TblInventoryTagUpdateOne {
	tituo.mutation.ClearInventory()
	return tituo
}

// ClearTag clears the "tag" edge to the TblTag entity.
func (tituo *TblInventoryTagUpdateOne) ClearTag() *TblInventoryTagUpdateOne {
	tituo.mutation.ClearTag()
	return tituo
}

// Where appends a list predicates to the TblInventoryTagUpdate builder.
func (tituo *TblInventoryTagUpdateOne) Where(ps ...predicate.TblInventoryTag) *TblInventoryTagUpdateOne {
	tituo.mutation.Where(ps...)
	return tituo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tituo *TblInventoryTagUpdateOne) Select(field string, fields ...string) *TblInventoryTagUpdateOne {
	tituo.fields = append([]string{field}, fields...)
	return tituo
}

// Save executes the query and returns the updated TblInventoryTag entity.
func (tituo *TblInventoryTagUpdateOne) Save(ctx context.Context) (*TblInventoryTag, error) {
	return withHooks(ctx, tituo.sqlSave, tituo.mutation, tituo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tituo *TblInventoryTagUpdateOne) SaveX(ctx context.Context) *TblInventoryTag {
	node, err := tituo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tituo *TblInventoryTagUpdateOne) Exec(ctx context.Context) error {
	_, err := tituo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tituo *TblInventoryTagUpdateOne) ExecX(ctx context.Context) {
	if err := tituo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tituo *TblInventoryTagUpdateOne) sqlSave(ctx context.Context) (_node *TblInventoryTag, err error) {
	_spec := sqlgraph.NewUpdateSpec(tblinventorytag.Table, tblinventorytag.Columns, sqlgraph.NewFieldSpec(tblinventorytag.FieldID, field.TypeString))
	id, ok := tituo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgen: missing "TblInventoryTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tituo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblinventorytag.FieldID)
		for _, f := range fields {
			if !tblinventorytag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != tblinventorytag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tituo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tituo.mutation.CreatedAt(); ok {
		_spec.SetField(tblinventorytag.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tituo.mutation.UpdatedAt(); ok {
		_spec.SetField(tblinventorytag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tituo.mutation.DeletedAt(); ok {
		_spec.SetField(tblinventorytag.FieldDeletedAt, field.TypeTime, value)
	}
	if tituo.mutation.DeletedAtCleared() {
		_spec.ClearField(tblinventorytag.FieldDeletedAt, field.TypeTime)
	}
	if tituo.mutation.InventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.InventoryTable,
			Columns: []string{tblinventorytag.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tituo.mutation.InventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.InventoryTable,
			Columns: []string{tblinventorytag.InventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tituo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.TagTable,
			Columns: []string{tblinventorytag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbltag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tituo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tblinventorytag.TagTable,
			Columns: []string{tblinventorytag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbltag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TblInventoryTag{config: tituo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tituo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblinventorytag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tituo.mutation.done = true
	return _node, nil
}
