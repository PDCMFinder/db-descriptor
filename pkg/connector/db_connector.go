// Package connector contains structs and interfaces to connect to the database and get the descriptions.
// The main piece is the [DBConnector] interface that defines what specific connectors should implement.
// Currently a specific implementation is provided (postgres) but others could be added in the future.
package connector

import "database/sql"

/*
An interface defining methods needed to get the descriptions from a database.

DatabaseDescription contains a slice of `Schema`.
*/
type DBConnector interface {
	// The database type of the connector. Example: postgres
	GetDatabaseTypeName() string
	// Gets a connection to the database after some credentials are provided
	GetConnection() (*sql.DB, error)
	/*
		A SQL query that brings the information for entities. Implementations are expected to provide the following columns:
		- table_schema (Schema of the entity)
		- table_name   (Entity name)
		- table_type   (Entity type: view/table)
		- comment      (Entity comment)

	*/
	GetEntitiesQueryStatement() string
	/*
		A SQL query that brings the information for columns. Implementations are expected to provide the following columns:
		- table_schema (Schema of the entity)
		- table_name   (Entity name)
		- column_name  (The name of the column)
		- data_type    (Data type of the column)
		- comment      (Column comment)

	*/
	GetColumnsQueryStatement() string
	/*
		A SQL query that brings relations between entities (FKs). Implementations are expected to provide the following columns:
		- table_schema 			(Schema of the entity)
		- table_name   			(Entity name)
		- column_name  			(The name of the column)
		- constraint_name   	(The name of the fk)
		- foreign_table_schema	(The name of the schema of the referenced table)
		- foreign_table_name	(The name of the referenced table)
		- foreign_column_name	(The name of the pk column in the referenced table)

	*/
	GetRelationsQueryStatement() string
}
