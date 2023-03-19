package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/JojoMakesGames/steam-graphql/dataloaders"
	"github.com/JojoMakesGames/steam-graphql/graph/model"
)

// Publishers is the resolver for the publishers field.
func (r *gameResolver) Publishers(ctx context.Context, obj *model.Game) ([]*model.Company, error) {
	publishers, err := dataloaders.GetCompanies(ctx, obj.PublisherIDs)
	if err != nil {
		return nil, err
	}
	return publishers, nil
}

// Developers is the resolver for the developers field.
func (r *gameResolver) Developers(ctx context.Context, obj *model.Game) ([]*model.Company, error) {
	developers, err := dataloaders.GetCompanies(ctx, obj.DeveloperIDs)
	if err != nil {
		return nil, err
	}
	return developers, nil
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	returnGames, err := r.GameService.GetGames()
	if err != nil {
		return nil, err
	}
	return returnGames, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type gameResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
