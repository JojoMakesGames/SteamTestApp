package helpers

import (
	"server/models"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GameService struct {
	Driver neo4j.DriverWithContext
}

func (gs GameService) GetGames() ([]*models.Game, error) {
	// session := gs.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	// defer session.Close()

	return nil, nil
}
