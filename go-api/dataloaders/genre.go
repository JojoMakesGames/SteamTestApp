package dataloaders

import (
	"context"
	"fmt"
	"strings"

	"github.com/JojoMakesGames/steam-graphql/graph/model"

	"github.com/graph-gophers/dataloader"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GenreReader struct {
	driver neo4j.Driver
}

func (g *GenreReader) GetGenres(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	genreIDs := make([]string, len(keys))
	for ix, key := range keys {
		genreIDs[ix] = fmt.Sprintf("\"%s\"", key.String())
	}
	query_string := fmt.Sprintf("MATCH (genre:Genre) WHERE elementId(genre) IN [%v] RETURN genre", strings.Join(genreIDs[:], ", "))
	session := g.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run(query_string, nil)
	if err != nil {
		return nil
	}

	genreById := map[string]*model.Genre{}
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("genre")
		g := node.(neo4j.Node)
		genre := &model.Genre{
			ID:   g.ElementId,
			Name: g.Props["name"].(string),
		}
		genreById[genre.ID] = genre
	}

	output := make([]*dataloader.Result, len(keys))
	for index, genreKey := range keys {
		company, ok := genreById[genreKey.String()]
		if !ok {
			output[index] = &dataloader.Result{Error: fmt.Errorf("no genre with id %s", genreKey)}
			continue
		} else {
			output[index] = &dataloader.Result{Data: company}
		}
	}
	fmt.Println(output)

	return output
}

func GetGenre(ctx context.Context, genreId string) (*model.Genre, error) {
	loaders := For(ctx)
	thunk := loaders.GenreLoader.Load(ctx, dataloader.StringKey(genreId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Genre), nil
}

func GetGenres(ctx context.Context, genreIds []string) ([]*model.Genre, error) {
	loaders := For(ctx)
	thunk := loaders.GenreLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(genreIds))
	result, _ := thunk()
	if result == nil {
		return make([]*model.Genre, 0), nil
	}
	genres := make([]*model.Genre, len(result))
	for ix, genre := range result {
		if genre == nil {
			continue
		}
		genres[ix] = genre.(*model.Genre)
	}

	return genres, nil
}
