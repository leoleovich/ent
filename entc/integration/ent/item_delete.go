// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// ItemDelete is the builder for deleting a Item entity.
type ItemDelete struct {
	config
	hooks      []Hook
	mutation   *ItemMutation
	predicates []predicate.Item
}

// Where adds a new predicate to the delete builder.
func (id *ItemDelete) Where(ps ...predicate.Item) *ItemDelete {
	id.predicates = append(id.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *ItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			return affected, err
		})
		for i := len(id.hooks); i > 0; i-- {
			mut = id.hooks[i-1](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *ItemDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *ItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	if ps := id.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// ItemDeleteOne is the builder for deleting a single Item entity.
type ItemDeleteOne struct {
	id *ItemDelete
}

// Exec executes the deletion query.
func (ido *ItemDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{item.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *ItemDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
