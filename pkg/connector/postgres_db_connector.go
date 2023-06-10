package connector

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

/*
A Postgres implementation of [DBConnector].

Implement methods to connect to a postgres database and obtain information about its tables, views, and columns.
*/
type PostgresDBConnector struct {
	Input Input
}

func (dbConnector PostgresDBConnector) GetDatabaseTypeName() string {
	return "Postgres"
}

func (dbConnector PostgresDBConnector) GetConnection() (*sql.DB, error) {
	postgresqlDbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnector.Input.Host,
		dbConnector.Input.Port,
		dbConnector.Input.User,
		dbConnector.Input.Password,
		dbConnector.Input.Name)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err == nil {
		err = db.Ping()
	}
	return db, err
}

func (dbConnector PostgresDBConnector) GetEntitiesQueryStatement() string {
	queryTemplate :=
		`SELECT 
		table_schema, table_name, table_type,
		COALESCE(obj_description(
		('"' || table_schema || '"."' || table_name || '"')::regclass,
		'pg_class'), '') AS comment
    FROM 
		information_schema.tables WHERE table_schema in ([SCHEMAS])`

	query := strings.Replace(queryTemplate, "[SCHEMAS]", getFormattedSchemaList(dbConnector.Input.Schemas), -1)
	return query
}

func (dbConnector PostgresDBConnector) GetColumnsQueryStatement() string {
	queryTemplate :=
		`SELECT 
		isc.table_schema,
		isc.table_name,
		isc.column_name,
		isc.data_type,
		COALESCE(pg_catalog.col_description(format('%s.%s',isc.table_schema,isc.table_name)::regclass::oid,isc.ordinal_position),'') as column_description
	FROM
		information_schema.columns isc WHERE table_schema in ([SCHEMAS])
		order by isc.table_name,isc.column_name`

	query := strings.Replace(queryTemplate, "[SCHEMAS]", getFormattedSchemaList(dbConnector.Input.Schemas), -1)
	return query
}
