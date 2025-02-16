package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/miladjlz/golang-graphql-gorm-postgresql/db"
	"github.com/miladjlz/golang-graphql-gorm-postgresql/graph/model"
)

var database, err = db.NewPostgresUserStore()

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {

	if err != nil {
		return nil, err
	}
	return database.InsertUser(input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UpdateUserInput) (*model.User, error) {
	if err != nil {
		return nil, err
	}
	return database.UpdateUser(id, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.DeleteUserResponse, error) {
	if err != nil {
		return nil, err
	}
	return database.DeleteUser(id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	if err != nil {
		return nil, err
	}
	return database.GetUsers()
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	if err != nil {
		return nil, err
	}
	return database.GetUser(id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
