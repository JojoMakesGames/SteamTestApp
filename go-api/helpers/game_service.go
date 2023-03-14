package helpers

import (
	"server/graph"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GameService struct {
	Driver neo4j.DriverWithContext
}

func (gs GameService) GetGames() ([]*Game, error) {
	session := gs.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	return returnGames, nil
}
