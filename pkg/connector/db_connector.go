package connector

import "database/sql"

type DBConnector interface {
	GetDatabaseTypeName() string
	GetConnection() (*sql.DB, error)
	GetEntitiesQueryStatement() string
	GetColumnsQueryStatement() string
}
