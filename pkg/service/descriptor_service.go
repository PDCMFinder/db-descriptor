// Package service contains the entry point of the logic to retrieve descriptions from a database.
package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/PDCMFinder/db-descriptor/internal/extractor"
	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/model"
)

// Returns a `model.DatabaseDescription` object with the description of a database.
func GetDbDescription(input connector.Input) model.DatabaseDescription {
	dbConnector, err := getDBConnector(input)
	if err != nil {
		log.Fatal(err)
	}
	dbDescriptionExtractor := extractor.New(dbConnector)
	databaseDescription := dbDescriptionExtractor.ExtractDescription()
	return databaseDescription
}

// Helper function to get the appropiate DBConnector implementation. Really simple logic as only one DBConnector
// is implemented. This could be much more sophisticated, following a plugin-like approach.
func getDBConnector(input connector.Input) (connector.DBConnector, error) {
	var dbConnector connector.DBConnector
	switch input.Db {
	case "postgres":
		dbConnector = connector.PostgresDBConnector{Input: input}
	default:
		return nil, errors.New(fmt.Sprintf("Database type [%s] not supported.", input.Db))
	}

	return dbConnector, nil
}
