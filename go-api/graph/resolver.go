package graph

import (
	"github.com/JojoMakesGames/steam-graphql/helpers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GameService helpers.GameService
}
