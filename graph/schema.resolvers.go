package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/site-tech/jaw-platform/cmd"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph/model"
)

// Jaw is the resolver for the jaw field.
func (r *queryResolver) Jaw(ctx context.Context) (*ent.User, error) {
	go cmd.Run()
	return r.client.User.Query().First(ctx)
}

// DbConnection is the resolver for the dbConnection field.
func (r *queryResolver) DbConnection(ctx context.Context, cred *model.DBConnection) (int, error) {
	fmt.Println("credentials: ", cred)
	return 200, nil
}
