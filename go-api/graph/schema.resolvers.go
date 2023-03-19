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

// ReleaseDate is the resolver for the release_date field.
func (r *gameResolver) ReleaseDate(ctx context.Context, obj *model.Game) (string, error) {
	panic(fmt.Errorf("not implemented: ReleaseDate - release_date"))
}

// Publishers is the resolver for the publishers field.
func (r *gameResolver) Publishers(ctx context.Context, obj *model.Game) ([]*model.Company, error) {
	publisherIds := make([]string, len(obj.Published))
	for i, published := range obj.Published {
		publisherIds[i] = published.StartElementId
	}
	publishers, err := dataloaders.GetCompanies(ctx, publisherIds)
	if err != nil {
		return nil, err
	}
	return publishers, nil
}

// Developers is the resolver for the developers field.
func (r *gameResolver) Developers(ctx context.Context, obj *model.Game) ([]*model.Company, error) {
	developerIds := make([]string, len(obj.Developed))
	for i, developed := range obj.Developed {
		developerIds[i] = developed.StartElementId
	}
	developers, err := dataloaders.GetCompanies(ctx, developerIds)
	if err != nil {
		return nil, err
	}
	return developers, nil
}

// Genres is the resolver for the genres field.
func (r *gameResolver) Genres(ctx context.Context, obj *model.Game) ([]*model.Genre, error) {
	genreIds := make([]string, len(obj.GameTypes))
	for i, gameType := range obj.GameTypes {
		genreIds[i] = gameType.EndElementId
	}
	genres, err := dataloaders.GetGenres(ctx, genreIds)
	if err != nil {
		return nil, err
	}
	return genres, nil
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
