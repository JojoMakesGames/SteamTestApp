package dataloaders

import (
	"context"
	"fmt"
	"strings"

	"github.com/JojoMakesGames/steam-graphql/graph/model"

	"github.com/graph-gophers/dataloader"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CompanyReader struct {
	driver neo4j.Driver
}

func (c *CompanyReader) GetCompanies(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	companyIDs := make([]string, len(keys))
	for ix, key := range keys {
		companyIDs[ix] = fmt.Sprintf("\"%s\"", key.String())
	}
	query_string := fmt.Sprintf("MATCH (company:Company) WHERE elementId(company) IN [%v] RETURN company", strings.Join(companyIDs[:], ", "))
	session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run(query_string, nil)
	if err != nil {
		return nil
	}

	companyById := map[string]*model.Company{}
	var record *neo4j.Record
	for result.NextRecord(&record) {
		node, _ := record.Get("company")
		c := node.(neo4j.Node)
		company := &model.Company{
			ID:   c.ElementId,
			Name: c.Props["name"].(string),
		}
		companyById[company.ID] = company
	}

	output := make([]*dataloader.Result, len(keys))
	for index, companyKey := range keys {
		company, ok := companyById[companyKey.String()]
		if !ok {
			output[index] = &dataloader.Result{Error: fmt.Errorf("no company with id %s", companyKey)}
			continue
		} else {
			output[index] = &dataloader.Result{Data: company}
		}
	}

	return output
}

func GetCompany(ctx context.Context, companyID string) (*model.Company, error) {
	loaders := For(ctx)
	thunk := loaders.CompanyLoader.Load(ctx, dataloader.StringKey(companyID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Company), nil
}

func GetCompanies(ctx context.Context, companyIDs []string) ([]*model.Company, error) {
	loaders := For(ctx)
	thunk := loaders.CompanyLoader.LoadMany(ctx, dataloader.NewKeysFromStrings(companyIDs))
	result, _ := thunk()
	if result == nil {
		return make([]*model.Company, 0), nil
	}
	companies := make([]*model.Company, len(result))
	for ix, company := range result {
		if company == nil {
			continue
		}
		companies[ix] = company.(*model.Company)
	}

	return companies, nil
}
