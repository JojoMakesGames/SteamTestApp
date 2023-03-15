package helpers

import (
	"server/graph/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GameService struct {
	Driver neo4j.Driver
}

func (gs GameService) GetGames() ([]*model.Game, error) {
	session := gs.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (game:Game) RETURN a", nil)
	if err != nil {
		return nil, err
	}
	returnGames := make([]*model.Game, 0)
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("game")
		game := node.(neo4j.Node)
		returnGames = append(returnGames, &model.Game{ID: game.ElementId, Name: game.Props["name"].(string)})
	}

	return returnGames, nil
}
