// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/realHoangHai/go-assmin/internal/ent/predicate"
	"github.com/realHoangHai/go-assmin/internal/ent/session"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SessionUpdate) SetUserID(u uuid.UUID) *SessionUpdate {
	su.mutation.SetUserID(u)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(u *uuid.UUID) *SessionUpdate {
	if u != nil {
		su.SetUserID(*u)
	}
	return su
}

// ClearUserID clears the value of the "user_id" field.
func (su *SessionUpdate) ClearUserID() *SessionUpdate {
	su.mutation.ClearUserID()
	return su
}

// SetRefreshToken sets the "refresh_token" field.
func (su *SessionUpdate) SetRefreshToken(s string) *SessionUpdate {
	su.mutation.SetRefreshToken(s)
	return su
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (su *SessionUpdate) SetNillableRefreshToken(s *string) *SessionUpdate {
	if s != nil {
		su.SetRefreshToken(*s)
	}
	return su
}

// ClearRefreshToken clears the value of the "refresh_token" field.
func (su *SessionUpdate) ClearRefreshToken() *SessionUpdate {
	su.mutation.ClearRefreshToken()
	return su
}

// SetUserAgent sets the "user_agent" field.
func (su *SessionUpdate) SetUserAgent(s string) *SessionUpdate {
	su.mutation.SetUserAgent(s)
	return su
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserAgent(s *string) *SessionUpdate {
	if s != nil {
		su.SetUserAgent(*s)
	}
	return su
}

// ClearUserAgent clears the value of the "user_agent" field.
func (su *SessionUpdate) ClearUserAgent() *SessionUpdate {
	su.mutation.ClearUserAgent()
	return su
}

// SetClientIP sets the "client_ip" field.
func (su *SessionUpdate) SetClientIP(s string) *SessionUpdate {
	su.mutation.SetClientIP(s)
	return su
}

// SetNillableClientIP sets the "client_ip" field if the given value is not nil.
func (su *SessionUpdate) SetNillableClientIP(s *string) *SessionUpdate {
	if s != nil {
		su.SetClientIP(*s)
	}
	return su
}

// ClearClientIP clears the value of the "client_ip" field.
func (su *SessionUpdate) ClearClientIP() *SessionUpdate {
	su.mutation.ClearClientIP()
	return su
}

// SetIsBlocked sets the "is_blocked" field.
func (su *SessionUpdate) SetIsBlocked(b bool) *SessionUpdate {
	su.mutation.SetIsBlocked(b)
	return su
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (su *SessionUpdate) SetNillableIsBlocked(b *bool) *SessionUpdate {
	if b != nil {
		su.SetIsBlocked(*b)
	}
	return su
}

// ClearIsBlocked clears the value of the "is_blocked" field.
func (su *SessionUpdate) ClearIsBlocked() *SessionUpdate {
	su.mutation.ClearIsBlocked()
	return su
}

// SetExpireTime sets the "expire_time" field.
func (su *SessionUpdate) SetExpireTime(t time.Time) *SessionUpdate {
	su.mutation.SetExpireTime(t)
	return su
}

// SetNillableExpireTime sets the "expire_time" field if the given value is not nil.
func (su *SessionUpdate) SetNillableExpireTime(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetExpireTime(*t)
	}
	return su
}

// ClearExpireTime clears the value of the "expire_time" field.
func (su *SessionUpdate) ClearExpireTime() *SessionUpdate {
	su.mutation.ClearExpireTime()
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SessionUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := session.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SessionUpdate) check() error {
	if v, ok := su.mutation.RefreshToken(); ok {
		if err := session.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "Session.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(session.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeUUID, value)
	}
	if su.mutation.UserIDCleared() {
		_spec.ClearField(session.FieldUserID, field.TypeUUID)
	}
	if value, ok := su.mutation.RefreshToken(); ok {
		_spec.SetField(session.FieldRefreshToken, field.TypeString, value)
	}
	if su.mutation.RefreshTokenCleared() {
		_spec.ClearField(session.FieldRefreshToken, field.TypeString)
	}
	if value, ok := su.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if su.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := su.mutation.ClientIP(); ok {
		_spec.SetField(session.FieldClientIP, field.TypeString, value)
	}
	if su.mutation.ClientIPCleared() {
		_spec.ClearField(session.FieldClientIP, field.TypeString)
	}
	if value, ok := su.mutation.IsBlocked(); ok {
		_spec.SetField(session.FieldIsBlocked, field.TypeBool, value)
	}
	if su.mutation.IsBlockedCleared() {
		_spec.ClearField(session.FieldIsBlocked, field.TypeBool)
	}
	if value, ok := su.mutation.ExpireTime(); ok {
		_spec.SetField(session.FieldExpireTime, field.TypeTime, value)
	}
	if su.mutation.ExpireTimeCleared() {
		_spec.ClearField(session.FieldExpireTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetUserID sets the "user_id" field.
func (suo *SessionUpdateOne) SetUserID(u uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetUserID(u)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(u *uuid.UUID) *SessionUpdateOne {
	if u != nil {
		suo.SetUserID(*u)
	}
	return suo
}

// ClearUserID clears the value of the "user_id" field.
func (suo *SessionUpdateOne) ClearUserID() *SessionUpdateOne {
	suo.mutation.ClearUserID()
	return suo
}

// SetRefreshToken sets the "refresh_token" field.
func (suo *SessionUpdateOne) SetRefreshToken(s string) *SessionUpdateOne {
	suo.mutation.SetRefreshToken(s)
	return suo
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableRefreshToken(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetRefreshToken(*s)
	}
	return suo
}

// ClearRefreshToken clears the value of the "refresh_token" field.
func (suo *SessionUpdateOne) ClearRefreshToken() *SessionUpdateOne {
	suo.mutation.ClearRefreshToken()
	return suo
}

// SetUserAgent sets the "user_agent" field.
func (suo *SessionUpdateOne) SetUserAgent(s string) *SessionUpdateOne {
	suo.mutation.SetUserAgent(s)
	return suo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserAgent(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUserAgent(*s)
	}
	return suo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (suo *SessionUpdateOne) ClearUserAgent() *SessionUpdateOne {
	suo.mutation.ClearUserAgent()
	return suo
}

// SetClientIP sets the "client_ip" field.
func (suo *SessionUpdateOne) SetClientIP(s string) *SessionUpdateOne {
	suo.mutation.SetClientIP(s)
	return suo
}

// SetNillableClientIP sets the "client_ip" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableClientIP(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetClientIP(*s)
	}
	return suo
}

// ClearClientIP clears the value of the "client_ip" field.
func (suo *SessionUpdateOne) ClearClientIP() *SessionUpdateOne {
	suo.mutation.ClearClientIP()
	return suo
}

// SetIsBlocked sets the "is_blocked" field.
func (suo *SessionUpdateOne) SetIsBlocked(b bool) *SessionUpdateOne {
	suo.mutation.SetIsBlocked(b)
	return suo
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableIsBlocked(b *bool) *SessionUpdateOne {
	if b != nil {
		suo.SetIsBlocked(*b)
	}
	return suo
}

// ClearIsBlocked clears the value of the "is_blocked" field.
func (suo *SessionUpdateOne) ClearIsBlocked() *SessionUpdateOne {
	suo.mutation.ClearIsBlocked()
	return suo
}

// SetExpireTime sets the "expire_time" field.
func (suo *SessionUpdateOne) SetExpireTime(t time.Time) *SessionUpdateOne {
	suo.mutation.SetExpireTime(t)
	return suo
}

// SetNillableExpireTime sets the "expire_time" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableExpireTime(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetExpireTime(*t)
	}
	return suo
}

// ClearExpireTime clears the value of the "expire_time" field.
func (suo *SessionUpdateOne) ClearExpireTime() *SessionUpdateOne {
	suo.mutation.ClearExpireTime()
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SessionUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := session.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SessionUpdateOne) check() error {
	if v, ok := suo.mutation.RefreshToken(); ok {
		if err := session.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "Session.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(session.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UserID(); ok {
		_spec.SetField(session.FieldUserID, field.TypeUUID, value)
	}
	if suo.mutation.UserIDCleared() {
		_spec.ClearField(session.FieldUserID, field.TypeUUID)
	}
	if value, ok := suo.mutation.RefreshToken(); ok {
		_spec.SetField(session.FieldRefreshToken, field.TypeString, value)
	}
	if suo.mutation.RefreshTokenCleared() {
		_spec.ClearField(session.FieldRefreshToken, field.TypeString)
	}
	if value, ok := suo.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if suo.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := suo.mutation.ClientIP(); ok {
		_spec.SetField(session.FieldClientIP, field.TypeString, value)
	}
	if suo.mutation.ClientIPCleared() {
		_spec.ClearField(session.FieldClientIP, field.TypeString)
	}
	if value, ok := suo.mutation.IsBlocked(); ok {
		_spec.SetField(session.FieldIsBlocked, field.TypeBool, value)
	}
	if suo.mutation.IsBlockedCleared() {
		_spec.ClearField(session.FieldIsBlocked, field.TypeBool)
	}
	if value, ok := suo.mutation.ExpireTime(); ok {
		_spec.SetField(session.FieldExpireTime, field.TypeTime, value)
	}
	if suo.mutation.ExpireTimeCleared() {
		_spec.ClearField(session.FieldExpireTime, field.TypeTime)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
