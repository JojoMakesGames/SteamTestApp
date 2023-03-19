package model

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type Game struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Relationships
	Published []*neo4j.Relationship `json:"published"`
	Developed []*neo4j.Relationship `json:"developed"`
	GameTypes []*neo4j.Relationship `json:"genres"`
}
