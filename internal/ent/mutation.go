// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

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
	create_time   *time.Time
	update_time   *time.Time
	status        *int
	addstatus     *int
	del_status    *int
	adddel_status *int
	user_name     *string
	true_name     *string
	mobile        *string
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

// SetCreateTime sets the "create_time" field.
func (m *AdminMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *AdminMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *AdminMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *AdminMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *AdminMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *AdminMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetStatus sets the "status" field.
func (m *AdminMutation) SetStatus(i int) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *AdminMutation) Status() (r int, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldStatus(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds i to the "status" field.
func (m *AdminMutation) AddStatus(i int) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *AdminMutation) AddedStatus() (r int, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *AdminMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// SetDelStatus sets the "del_status" field.
func (m *AdminMutation) SetDelStatus(i int) {
	m.del_status = &i
	m.adddel_status = nil
}

// DelStatus returns the value of the "del_status" field in the mutation.
func (m *AdminMutation) DelStatus() (r int, exists bool) {
	v := m.del_status
	if v == nil {
		return
	}
	return *v, true
}

// OldDelStatus returns the old "del_status" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldDelStatus(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDelStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDelStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDelStatus: %w", err)
	}
	return oldValue.DelStatus, nil
}

// AddDelStatus adds i to the "del_status" field.
func (m *AdminMutation) AddDelStatus(i int) {
	if m.adddel_status != nil {
		*m.adddel_status += i
	} else {
		m.adddel_status = &i
	}
}

// AddedDelStatus returns the value that was added to the "del_status" field in this mutation.
func (m *AdminMutation) AddedDelStatus() (r int, exists bool) {
	v := m.adddel_status
	if v == nil {
		return
	}
	return *v, true
}

// ResetDelStatus resets all changes to the "del_status" field.
func (m *AdminMutation) ResetDelStatus() {
	m.del_status = nil
	m.adddel_status = nil
}

// SetUserName sets the "user_name" field.
func (m *AdminMutation) SetUserName(s string) {
	m.user_name = &s
}

// UserName returns the value of the "user_name" field in the mutation.
func (m *AdminMutation) UserName() (r string, exists bool) {
	v := m.user_name
	if v == nil {
		return
	}
	return *v, true
}

// OldUserName returns the old "user_name" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldUserName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserName: %w", err)
	}
	return oldValue.UserName, nil
}

// ResetUserName resets all changes to the "user_name" field.
func (m *AdminMutation) ResetUserName() {
	m.user_name = nil
}

// SetTrueName sets the "true_name" field.
func (m *AdminMutation) SetTrueName(s string) {
	m.true_name = &s
}

// TrueName returns the value of the "true_name" field in the mutation.
func (m *AdminMutation) TrueName() (r string, exists bool) {
	v := m.true_name
	if v == nil {
		return
	}
	return *v, true
}

// OldTrueName returns the old "true_name" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldTrueName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTrueName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTrueName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTrueName: %w", err)
	}
	return oldValue.TrueName, nil
}

// ResetTrueName resets all changes to the "true_name" field.
func (m *AdminMutation) ResetTrueName() {
	m.true_name = nil
}

// SetMobile sets the "mobile" field.
func (m *AdminMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *AdminMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the Admin entity.
// If the Admin object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AdminMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *AdminMutation) ResetMobile() {
	m.mobile = nil
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
	fields := make([]string, 0, 7)
	if m.create_time != nil {
		fields = append(fields, admin.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, admin.FieldUpdateTime)
	}
	if m.status != nil {
		fields = append(fields, admin.FieldStatus)
	}
	if m.del_status != nil {
		fields = append(fields, admin.FieldDelStatus)
	}
	if m.user_name != nil {
		fields = append(fields, admin.FieldUserName)
	}
	if m.true_name != nil {
		fields = append(fields, admin.FieldTrueName)
	}
	if m.mobile != nil {
		fields = append(fields, admin.FieldMobile)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AdminMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case admin.FieldCreateTime:
		return m.CreateTime()
	case admin.FieldUpdateTime:
		return m.UpdateTime()
	case admin.FieldStatus:
		return m.Status()
	case admin.FieldDelStatus:
		return m.DelStatus()
	case admin.FieldUserName:
		return m.UserName()
	case admin.FieldTrueName:
		return m.TrueName()
	case admin.FieldMobile:
		return m.Mobile()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AdminMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case admin.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case admin.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case admin.FieldStatus:
		return m.OldStatus(ctx)
	case admin.FieldDelStatus:
		return m.OldDelStatus(ctx)
	case admin.FieldUserName:
		return m.OldUserName(ctx)
	case admin.FieldTrueName:
		return m.OldTrueName(ctx)
	case admin.FieldMobile:
		return m.OldMobile(ctx)
	}
	return nil, fmt.Errorf("unknown Admin field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) SetField(name string, value ent.Value) error {
	switch name {
	case admin.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case admin.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case admin.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case admin.FieldDelStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDelStatus(v)
		return nil
	case admin.FieldUserName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserName(v)
		return nil
	case admin.FieldTrueName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTrueName(v)
		return nil
	case admin.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	}
	return fmt.Errorf("unknown Admin field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AdminMutation) AddedFields() []string {
	var fields []string
	if m.addstatus != nil {
		fields = append(fields, admin.FieldStatus)
	}
	if m.adddel_status != nil {
		fields = append(fields, admin.FieldDelStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AdminMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case admin.FieldStatus:
		return m.AddedStatus()
	case admin.FieldDelStatus:
		return m.AddedDelStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AdminMutation) AddField(name string, value ent.Value) error {
	switch name {
	case admin.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	case admin.FieldDelStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDelStatus(v)
		return nil
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
	case admin.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case admin.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case admin.FieldStatus:
		m.ResetStatus()
		return nil
	case admin.FieldDelStatus:
		m.ResetDelStatus()
		return nil
	case admin.FieldUserName:
		m.ResetUserName()
		return nil
	case admin.FieldTrueName:
		m.ResetTrueName()
		return nil
	case admin.FieldMobile:
		m.ResetMobile()
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
