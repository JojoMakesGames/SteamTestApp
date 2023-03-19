package model

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Relationships
	FriendsWith []*neo4j.Relationship `json:"friendsWith"`
	Owns        []*neo4j.Relationship `json:"owns"`
}
