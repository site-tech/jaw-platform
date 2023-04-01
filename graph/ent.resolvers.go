package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph/generated"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.AccountWhereInput) (*ent.AccountConnection, error) {
	panic(fmt.Errorf("not implemented: Accounts - accounts"))
}

// Tennants is the resolver for the tennants field.
func (r *queryResolver) Tennants(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TennantWhereInput) (*ent.TennantConnection, error) {
	panic(fmt.Errorf("not implemented: Tennants - tennants"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
