package graph

import (
	"context"
	"fmt"
	"graphql-go/graph/model"
	"graphql-go/internal/links"
	"strconv"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	//var user model.User

	link.Title = input.Title
	link.Address = input.Address

	linkId := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkId, 10), Title: link.Title, Address: link.Address}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()

	for _, link := range dbLinks {
		resLinks = append(resLinks, &model.Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
			User:    nil,
		})
	}

	return resLinks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
