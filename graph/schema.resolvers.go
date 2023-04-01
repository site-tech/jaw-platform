package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/site-tech/jaw-platform/ent"
)

// Jaw is the resolver for the jaw field.
func (r *queryResolver) Jaw(ctx context.Context) (*ent.User, error) {
	return r.client.User.Query().First(ctx)
}
