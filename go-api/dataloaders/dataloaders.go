package dataloaders

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	CompanyLoader   *dataloader.Loader
	PublisherLoader *dataloader.Loader
	DeveloperLoader *dataloader.Loader
	GenreLoader     *dataloader.Loader
	GameLoader      *dataloader.Loader
}

func CreateLoaders(driver *neo4j.Driver) *Loaders {
	companyReader := &CompanyReader{driver: *driver}
	genreReader := &GenreReader{driver: *driver}
	gameReader := &GameReader{driver: *driver}
	loaders := &Loaders{
		CompanyLoader: dataloader.NewBatchedLoader(companyReader.GetCompanies),
		GenreLoader:   dataloader.NewBatchedLoader(genreReader.GetGenres),
		GameLoader:    dataloader.NewBatchedLoader(gameReader.GetGames),
	}

	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
