// Package extractor contains the logic to extract descriptions from a database.
package extractor

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/model"
)

/*
A struct to hold a [connector/DBConnector] and use it to query the database.

Tje `dBConnector` property is a implementation of [connector/DBConnector], specific to a database type (like postgres).
*/
type dbDescriptionExtractor struct {
	dBConnector connector.DBConnector
}

// Returns an instance of [dbDescriptionExtractor] after initializing it with a [connector/DBConnector].
func New(dBConnector connector.DBConnector) dbDescriptionExtractor {
	instance := dbDescriptionExtractor{dBConnector}
	return instance
}

// The description of the database is a list of `schema` objects. Each schema has the description of its entities and columns
func (d dbDescriptionExtractor) ExtractDescription() model.DatabaseDescription {
	log.Println("Init database description extraction. Database type:", d.dBConnector.GetDatabaseTypeName())

	// Init connection to the database
	db, err := d.dBConnector.GetConnection()
	validateConnection(db, err)

	// 2-dimensional map with schema name --> entity name --> entity
	dataMap := make(map[string]map[string]model.Entity)
	// Add descriptions of entities
	populateEntities(dataMap, d.dBConnector.GetEntitiesQueryStatement(), db)
	// Add descriptions of columns
	populateColumns(dataMap, d.dBConnector.GetColumnsQueryStatement(), db)

	defer db.Close()

	schemas := buildSchemeList(dataMap)
	return model.DatabaseDescription{Schemas: schemas}
}

// Populates `dataMap` with the database entities information
func populateEntities(dataMap map[string]map[string]model.Entity, queryStatement string, db *sql.DB) {
	entities, err := getEntitiesList(queryStatement, db)
	if err == nil {
		for _, e := range entities {
			// Check if the schema name already exists in the map. If not, then create the entry with the schema and an empty Entity map
			entityMap, schemaExists := dataMap[e.SchemaName]
			if !schemaExists {
				entityMap = make(map[string]model.Entity)
				dataMap[e.SchemaName] = entityMap
			}
			dataMap[e.SchemaName][e.Name] = e
		}
	}
}

// Populates `dataMap` with the database columns information
func populateColumns(dataMap map[string]map[string]model.Entity, queryStatement string, db *sql.DB) {
	columns, err := getColumnsList(queryStatement, db)
	if err == nil {
		for _, c := range columns {
			// The entity should exists already. If not, maybe an error could be thrown
			entity, entityExists := dataMap[c.SchemaName][c.EntityName]
			if entityExists {
				entity.Columns = append(entity.Columns, c)
				dataMap[c.SchemaName][c.EntityName] = entity
			}
		}
	}
}

// Executes the query to retrieve the entities and converts it to a list of `model.Entity`
func getEntitiesList(queryStatement string, db *sql.DB) ([]model.Entity, error) {
	rows, err := db.Query(queryStatement)
	if err == nil {
		return processEntityRows(rows), nil
	} else {
		return nil, err
	}
}

// Executes the query to retrieve the columns and converts it to a list of `model.Column`
func getColumnsList(queryStatement string, db *sql.DB) ([]model.Column, error) {
	rows, err := db.Query(queryStatement)
	if err == nil {
		return processColumnRows(rows), nil
	} else {
		return nil, err
	}
}

func validateConnection(db *sql.DB, err error) {
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("Could not connect to the database. Error: %s", err.Error())))
	}
}

// Converts the rows that contain the results of querying the entities in the database into a list of `model.Entity`
func processEntityRows(rows *sql.Rows) []model.Entity {
	entities := make([]model.Entity, 0)
	for rows.Next() {
		var entity_schema string
		var entity_name string
		var entity_type string
		var entity_comment string

		err := rows.Scan(&entity_schema, &entity_name, &entity_type, &entity_comment)
		if err != nil {
			panic(err)
		}

		schemaName := strings.ToLower(entity_schema)
		entityName := strings.ToLower(entity_name)
		entityType := processType(entity_type)
		entityComment := entity_comment

		var entity model.Entity = model.Entity{
			Name:       entityName,
			EntityType: entityType,
			Comment:    entityComment,
			SchemaName: schemaName}

		entities = append(entities, entity)
	}

	return entities
}

// Converts the rows that contain the results of querying the columns in the database into a list of `model.Column`
func processColumnRows(rows *sql.Rows) []model.Column {
	columns := make([]model.Column, 0)
	for rows.Next() {
		var entity_schema string
		var entity_name string
		var column_name string
		var data_type string
		var column_comment sql.NullString
		var is_primary_key sql.NullBool
		var is_foreign_key sql.NullBool

		err := rows.Scan(
			&entity_schema, &entity_name, &column_name, &data_type, &column_comment, &is_primary_key, &is_foreign_key)
		if err != nil {
			panic(err)
		}

		schemaName := strings.ToLower(entity_schema)
		entityName := strings.ToLower(entity_name)
		columnName := strings.ToLower(column_name)
		dataType := data_type

		var columnComment string
		if column_comment.Valid {
			columnComment = column_comment.String
		}

		var isPrimaryKey bool
		if is_primary_key.Valid {
			isPrimaryKey = is_primary_key.Bool
		}

		var isForeignKey bool
		if is_foreign_key.Valid {
			isForeignKey = is_foreign_key.Bool
		}
		var column model.Column = model.Column{
			SchemaName:   schemaName,
			EntityName:   entityName,
			Name:         columnName,
			DataType:     dataType,
			Comment:      columnComment,
			IsPrimaryKey: isPrimaryKey,
			IsForeignKey: isForeignKey}

		columns = append(columns, column)
	}

	return columns
}

func buildSchemeList(dataMap map[string]map[string]model.Entity) []model.Schema {
	var schemas = make([]model.Schema, 0)
	// Populate the list of schemas
	for schemaKey, entityMap := range dataMap {
		entities := make([]model.Entity, 0, len(entityMap))
		for _, value := range entityMap {
			entities = append(entities, value)
		}
		schema := model.Schema{Name: schemaKey, Entities: entities}
		schemas = append(schemas, schema)
	}
	return schemas
}

func processType(originalType string) string {
	entityType := strings.ToLower(originalType)
	if strings.Contains(entityType, "view") {
		entityType = "view"
	} else if strings.Contains(entityType, "table") {
		entityType = "table"
	}
	return entityType
}
