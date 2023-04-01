// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/site-tech/jaw-platform/ent/account"
	"github.com/site-tech/jaw-platform/ent/predicate"
	"github.com/site-tech/jaw-platform/ent/tennant"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
	TypeTennant = "Tennant"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	name            *string
	created_at      *time.Time
	clearedFields   map[string]struct{}
	tennants        map[uuid.UUID]struct{}
	removedtennants map[uuid.UUID]struct{}
	clearedtennants bool
	done            bool
	oldValue        func(context.Context) (*Account, error)
	predicates      []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id uuid.UUID) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Account entities.
func (m *AccountMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *AccountMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AccountMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AccountMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *AccountMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AccountMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AccountMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddTennantIDs adds the "tennants" edge to the Tennant entity by ids.
func (m *AccountMutation) AddTennantIDs(ids ...uuid.UUID) {
	if m.tennants == nil {
		m.tennants = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.tennants[ids[i]] = struct{}{}
	}
}

// ClearTennants clears the "tennants" edge to the Tennant entity.
func (m *AccountMutation) ClearTennants() {
	m.clearedtennants = true
}

// TennantsCleared reports if the "tennants" edge to the Tennant entity was cleared.
func (m *AccountMutation) TennantsCleared() bool {
	return m.clearedtennants
}

// RemoveTennantIDs removes the "tennants" edge to the Tennant entity by IDs.
func (m *AccountMutation) RemoveTennantIDs(ids ...uuid.UUID) {
	if m.removedtennants == nil {
		m.removedtennants = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.tennants, ids[i])
		m.removedtennants[ids[i]] = struct{}{}
	}
}

// RemovedTennants returns the removed IDs of the "tennants" edge to the Tennant entity.
func (m *AccountMutation) RemovedTennantsIDs() (ids []uuid.UUID) {
	for id := range m.removedtennants {
		ids = append(ids, id)
	}
	return
}

// TennantsIDs returns the "tennants" edge IDs in the mutation.
func (m *AccountMutation) TennantsIDs() (ids []uuid.UUID) {
	for id := range m.tennants {
		ids = append(ids, id)
	}
	return
}

// ResetTennants resets all changes to the "tennants" edge.
func (m *AccountMutation) ResetTennants() {
	m.tennants = nil
	m.clearedtennants = false
	m.removedtennants = nil
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, account.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, account.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldName:
		return m.Name()
	case account.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldName:
		return m.OldName(ctx)
	case account.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case account.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldName:
		m.ResetName()
		return nil
	case account.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tennants != nil {
		edges = append(edges, account.EdgeTennants)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case account.EdgeTennants:
		ids := make([]ent.Value, 0, len(m.tennants))
		for id := range m.tennants {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtennants != nil {
		edges = append(edges, account.EdgeTennants)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case account.EdgeTennants:
		ids := make([]ent.Value, 0, len(m.removedtennants))
		for id := range m.removedtennants {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtennants {
		edges = append(edges, account.EdgeTennants)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	switch name {
	case account.EdgeTennants:
		return m.clearedtennants
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	switch name {
	case account.EdgeTennants:
		m.ResetTennants()
		return nil
	}
	return fmt.Errorf("unknown Account edge %s", name)
}

// TennantMutation represents an operation that mutates the Tennant nodes in the graph.
type TennantMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	externalId     *string
	cloud          *string
	created_at     *time.Time
	clearedFields  map[string]struct{}
	account        *uuid.UUID
	clearedaccount bool
	done           bool
	oldValue       func(context.Context) (*Tennant, error)
	predicates     []predicate.Tennant
}

var _ ent.Mutation = (*TennantMutation)(nil)

// tennantOption allows management of the mutation configuration using functional options.
type tennantOption func(*TennantMutation)

// newTennantMutation creates new mutation for the Tennant entity.
func newTennantMutation(c config, op Op, opts ...tennantOption) *TennantMutation {
	m := &TennantMutation{
		config:        c,
		op:            op,
		typ:           TypeTennant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTennantID sets the ID field of the mutation.
func withTennantID(id uuid.UUID) tennantOption {
	return func(m *TennantMutation) {
		var (
			err   error
			once  sync.Once
			value *Tennant
		)
		m.oldValue = func(ctx context.Context) (*Tennant, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tennant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTennant sets the old Tennant of the mutation.
func withTennant(node *Tennant) tennantOption {
	return func(m *TennantMutation) {
		m.oldValue = func(context.Context) (*Tennant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TennantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TennantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Tennant entities.
func (m *TennantMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TennantMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TennantMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tennant.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetExternalId sets the "externalId" field.
func (m *TennantMutation) SetExternalId(s string) {
	m.externalId = &s
}

// ExternalId returns the value of the "externalId" field in the mutation.
func (m *TennantMutation) ExternalId() (r string, exists bool) {
	v := m.externalId
	if v == nil {
		return
	}
	return *v, true
}

// OldExternalId returns the old "externalId" field's value of the Tennant entity.
// If the Tennant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TennantMutation) OldExternalId(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExternalId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExternalId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExternalId: %w", err)
	}
	return oldValue.ExternalId, nil
}

// ResetExternalId resets all changes to the "externalId" field.
func (m *TennantMutation) ResetExternalId() {
	m.externalId = nil
}

// SetCloud sets the "cloud" field.
func (m *TennantMutation) SetCloud(s string) {
	m.cloud = &s
}

// Cloud returns the value of the "cloud" field in the mutation.
func (m *TennantMutation) Cloud() (r string, exists bool) {
	v := m.cloud
	if v == nil {
		return
	}
	return *v, true
}

// OldCloud returns the old "cloud" field's value of the Tennant entity.
// If the Tennant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TennantMutation) OldCloud(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCloud is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCloud requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCloud: %w", err)
	}
	return oldValue.Cloud, nil
}

// ResetCloud resets all changes to the "cloud" field.
func (m *TennantMutation) ResetCloud() {
	m.cloud = nil
}

// SetAccountID sets the "account_id" field.
func (m *TennantMutation) SetAccountID(u uuid.UUID) {
	m.account = &u
}

// AccountID returns the value of the "account_id" field in the mutation.
func (m *TennantMutation) AccountID() (r uuid.UUID, exists bool) {
	v := m.account
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountID returns the old "account_id" field's value of the Tennant entity.
// If the Tennant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TennantMutation) OldAccountID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountID: %w", err)
	}
	return oldValue.AccountID, nil
}

// ResetAccountID resets all changes to the "account_id" field.
func (m *TennantMutation) ResetAccountID() {
	m.account = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TennantMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TennantMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Tennant entity.
// If the Tennant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TennantMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TennantMutation) ResetCreatedAt() {
	m.created_at = nil
}

// ClearAccount clears the "account" edge to the Account entity.
func (m *TennantMutation) ClearAccount() {
	m.clearedaccount = true
}

// AccountCleared reports if the "account" edge to the Account entity was cleared.
func (m *TennantMutation) AccountCleared() bool {
	return m.clearedaccount
}

// AccountIDs returns the "account" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AccountID instead. It exists only for internal usage by the builders.
func (m *TennantMutation) AccountIDs() (ids []uuid.UUID) {
	if id := m.account; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAccount resets all changes to the "account" edge.
func (m *TennantMutation) ResetAccount() {
	m.account = nil
	m.clearedaccount = false
}

// Where appends a list predicates to the TennantMutation builder.
func (m *TennantMutation) Where(ps ...predicate.Tennant) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TennantMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Tennant).
func (m *TennantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TennantMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.externalId != nil {
		fields = append(fields, tennant.FieldExternalId)
	}
	if m.cloud != nil {
		fields = append(fields, tennant.FieldCloud)
	}
	if m.account != nil {
		fields = append(fields, tennant.FieldAccountID)
	}
	if m.created_at != nil {
		fields = append(fields, tennant.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TennantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tennant.FieldExternalId:
		return m.ExternalId()
	case tennant.FieldCloud:
		return m.Cloud()
	case tennant.FieldAccountID:
		return m.AccountID()
	case tennant.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TennantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tennant.FieldExternalId:
		return m.OldExternalId(ctx)
	case tennant.FieldCloud:
		return m.OldCloud(ctx)
	case tennant.FieldAccountID:
		return m.OldAccountID(ctx)
	case tennant.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Tennant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TennantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tennant.FieldExternalId:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExternalId(v)
		return nil
	case tennant.FieldCloud:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCloud(v)
		return nil
	case tennant.FieldAccountID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountID(v)
		return nil
	case tennant.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Tennant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TennantMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TennantMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TennantMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tennant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TennantMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TennantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TennantMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tennant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TennantMutation) ResetField(name string) error {
	switch name {
	case tennant.FieldExternalId:
		m.ResetExternalId()
		return nil
	case tennant.FieldCloud:
		m.ResetCloud()
		return nil
	case tennant.FieldAccountID:
		m.ResetAccountID()
		return nil
	case tennant.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Tennant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TennantMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.account != nil {
		edges = append(edges, tennant.EdgeAccount)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TennantMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tennant.EdgeAccount:
		if id := m.account; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TennantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TennantMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TennantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedaccount {
		edges = append(edges, tennant.EdgeAccount)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TennantMutation) EdgeCleared(name string) bool {
	switch name {
	case tennant.EdgeAccount:
		return m.clearedaccount
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TennantMutation) ClearEdge(name string) error {
	switch name {
	case tennant.EdgeAccount:
		m.ClearAccount()
		return nil
	}
	return fmt.Errorf("unknown Tennant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TennantMutation) ResetEdge(name string) error {
	switch name {
	case tennant.EdgeAccount:
		m.ResetAccount()
		return nil
	}
	return fmt.Errorf("unknown Tennant edge %s", name)
}