package dataloaders

import (
	"context"
	"fmt"
	"strings"

	"github.com/JojoMakesGames/steam-graphql/graph/model"

	"github.com/graph-gophers/dataloader"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GameReader struct {
	driver neo4j.Driver
}

func (g *GameReader) GetGames(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	gameIDs := make([]string, len(keys))
	for ix, key := range keys {
		gameIDs[ix] = fmt.Sprintf("\"%s\"", key.String())
	}
	query_string := fmt.Sprintf("MATCH (game:Game)-[rela]-() WHERE elementId(game) IN [%v] RETURN game, rela", strings.Join(gameIDs[:], ", "))
	session := g.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run(query_string, nil)
	if err != nil {
		return nil
	}

	gameById := map[string]*model.Game{}
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("game")
		g := node.(neo4j.Node)
		rel, _ := record.Get("rela")
		relationship := rel.(neo4j.Relationship)
		if gameById[g.ElementId] == nil {
			gameById[g.ElementId] = &model.Game{
				ID:   g.ElementId,
				Name: g.Props["name"].(string),
			}
		}
		switch relationship.Type {
		case "PUBLISHED":
			gameById[g.ElementId].Published = append(gameById[g.ElementId].Published, &relationship)
			break
		case "DEVELOPED":
			gameById[g.ElementId].Developed = append(gameById[g.ElementId].Developed, &relationship)
			break
		case "GENRE":
			gameById[g.ElementId].GameTypes = append(gameById[g.ElementId].GameTypes, &relationship)
			break
		}
	}

	output := make([]*dataloader.Result, len(keys))
	for index, gameKey := range keys {
		game, ok := gameById[gameKey.String()]
		if !ok {
			output[index] = &dataloader.Result{Error: fmt.Errorf("no genre with id %s", gameKey)}
			continue
		} else {
			output[index] = &dataloader.Result{Data: game}
		}
	}

	return output
}

func GetGame(ctx context.Context, gameId string) (*model.Game, error) {
	loaders := For(ctx)
	thunk := loaders.GameLoader.Load(ctx, dataloader.StringKey(gameId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Game), nil
}

func GetGames(ctx context.Context, gameIds []string) ([]*model.Game, error) {
	loaders := For(ctx)
	thunk := loaders.GameLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(gameIds))
	result, _ := thunk()
	if result == nil {
		return make([]*model.Game, 0), nil
	}
	games := make([]*model.Game, len(result))
	for ix, genre := range result {
		if genre == nil {
			continue
		}
		games[ix] = genre.(*model.Game)
	}

	return games, nil
}
