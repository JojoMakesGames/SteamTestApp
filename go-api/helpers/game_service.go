package helpers

import (
	"github.com/JojoMakesGames/steam-graphql/graph/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GameService struct {
	Driver neo4j.Driver
}

func (gs *GameService) GetGames() ([]*model.Game, error) {
	session := gs.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (game:Game)-[rela]-() RETURN game, rela", nil)
	if err != nil {
		return nil, err
	}
	returnGames := make([]*model.Game, 0)
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("game")
		rel, _ := record.Get("rela")
		relType := rel.(neo4j.Relationship).Type
		game := node.(neo4j.Node)
		relationship := rel.(neo4j.Relationship)
		if len(returnGames) > 0 && returnGames[len(returnGames)-1].ID == game.ElementId {
			switch relType {
			case "PUBLISHED":
				returnGames[len(returnGames)-1].Published = append(returnGames[len(returnGames)-1].Published, &relationship)
				break
			case "DEVELOPED":
				returnGames[len(returnGames)-1].Developed = append(returnGames[len(returnGames)-1].Developed, &relationship)
				break
			case "GAME_TYPE":
				returnGames[len(returnGames)-1].GameTypes = append(returnGames[len(returnGames)-1].GameTypes, &relationship)
			}
			continue
		}

		modelGame := &model.Game{
			ID:        game.ElementId,
			Name:      game.Props["name"].(string),
			Developed: make([]*neo4j.Relationship, 0),
			Published: make([]*neo4j.Relationship, 0),
			GameTypes: make([]*neo4j.Relationship, 0),
		}
		switch relType {
		case "PUBLISHED":
			modelGame.Published = append(modelGame.Published, &relationship)
			break
		case "DEVELOPED":
			modelGame.Developed = append(modelGame.Developed, &relationship)
			break
		case "GAME_TYPE":
			modelGame.GameTypes = append(modelGame.GameTypes, &relationship)
			break
		}

		returnGames = append(returnGames, modelGame)
	}

	return returnGames, nil
}
