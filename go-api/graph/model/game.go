package model

type Game struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	ReleaseDate  string   `json:"releaseDate"`
	PublisherIDs []string `json:"publisherIDs"`
	DeveloperIDs []string `json:"developerIDs"`
}
