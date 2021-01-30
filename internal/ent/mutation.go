// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/daymenu/goadmin/internal/ent/admin"
	"github.com/daymenu/goadmin/internal/ent/predicate"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAdmin = "Admin"
)

// AdminMutation represents an operation that mutates the Admin nodes in the graph.
type AdminMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Admin, error)
	predicates    []predicate.Admin
}

var _ ent.Mutation = (*AdminMutation)(nil)

// adminOption allows management of the mutation configuration using functional options.
type adminOption func(*AdminMutation)

// newAdminMutation creates new mutation for the Admin entity.
func newAdminMutation(c config, op Op, opts ...adminOption) *AdminMutation {
	m := &AdminMutation{
		config:        c,
		op:            op,
		typ:           TypeAdmin,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAdminID sets the ID field of the mutation.
func withAdminID(id int) adminOption {
	return func(m *AdminMutation) {
		var (
			err   error
			once  sync.Once
			value *Admin
		)
		m.oldValue = func(ctx context.Context) (*Admin, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Admin.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAdmin sets the old Admin of the mutation.
func withAdmin(node *Admin) adminOption {
	return func(m *AdminMutation) {
		m.oldValue = func(context.Context) (*Admin, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AdminMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AdminMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *AdminMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *AdminMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AdminMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AdminMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *AdminMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Admin).
func (m *AdminMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AdminMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, admin.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AdminMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case admin.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AdminMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case admin.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Admin field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) SetField(name string, value ent.Value) error {
	switch name {
	case admin.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Admin field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AdminMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AdminMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Admin numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AdminMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AdminMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AdminMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Admin nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AdminMutation) ResetField(name string) error {
	switch name {
	case admin.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Admin field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AdminMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AdminMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AdminMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AdminMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AdminMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AdminMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AdminMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Admin unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AdminMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Admin edge %s", name)
}
