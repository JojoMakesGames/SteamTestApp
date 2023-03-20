package helpers

import (
	"github.com/JojoMakesGames/steam-graphql/graph/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UserService struct {
	Driver neo4j.Driver
}

func (us *UserService) GetUsers() ([]*model.User, error) {
	session := us.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("MATCH (user:User)-[rela]-() RETURN user, rela", nil)
	if err != nil {
		return nil, err
	}
	returnUsers := make([]*model.User, 0)
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("user")
		rel, _ := record.Get("rela")
		relType := rel.(neo4j.Relationship).Type
		user := node.(neo4j.Node)
		relationship := rel.(neo4j.Relationship)
		if len(returnUsers) > 0 && returnUsers[len(returnUsers)-1].ID == user.ElementId {
			switch relType {
			case "FRIENDS_WITH":
				returnUsers[len(returnUsers)-1].FriendsWith = append(returnUsers[len(returnUsers)-1].FriendsWith, &relationship)
				break
			case "OWNS":
				returnUsers[len(returnUsers)-1].Owns = append(returnUsers[len(returnUsers)-1].Owns, &relationship)
				break
			}
			continue
		}

		modelUser := &model.User{
			ID:          user.ElementId,
			Name:        user.Props["name"].(string),
			FriendsWith: make([]*neo4j.Relationship, 0),
			Owns:        make([]*neo4j.Relationship, 0),
		}
		switch relType {
		case "FRIENDS_WITH":
			modelUser.FriendsWith = append(modelUser.FriendsWith, &relationship)
			break
		case "OWNS":
			modelUser.Owns = append(modelUser.Owns, &relationship)
			break
		}
		returnUsers = append(returnUsers, modelUser)

	}
	return returnUsers, nil
}
