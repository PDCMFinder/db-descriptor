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
		information_schema.tables WHERE table_schema in ([SCHEMAS])
		ORDER BY table_name`

	query := strings.Replace(queryTemplate, "[SCHEMAS]", getFormattedSchemaList(dbConnector.Input.Schemas), -1)
	return query
}

func (dbConnector PostgresDBConnector) GetColumnsQueryStatement() string {
	queryTemplate :=
		`SELECT
		ns.nspname AS schema_name,
		tbl.relname AS table_name,
		col.attname AS column_name,
		format_type(col.atttypid, col.atttypmod) AS column_type,
		col_description(tbl.oid, col.attnum) AS column_comment,
		(SELECT CASE WHEN con.conname IS NULL THEN FALSE ELSE TRUE END
		 FROM pg_constraint con
		 WHERE con.contype = 'p' AND con.conrelid = tbl.oid AND col.attnum = ANY(con.conkey)) AS is_primary_key,
		(SELECT CASE WHEN con.conname IS NULL THEN FALSE ELSE TRUE END
		 FROM pg_constraint con
		 WHERE con.contype = 'f' AND con.conrelid = tbl.oid AND col.attnum = ANY(con.conkey)) AS is_foreign_key
	FROM
		pg_namespace ns
		JOIN pg_class tbl ON tbl.relnamespace = ns.oid
		JOIN pg_attribute col ON col.attrelid = tbl.oid
	WHERE
		ns.nspname in ([SCHEMAS])
		AND tbl.relkind IN ('r', 'v')
		AND col.attnum > 0 -- Exclude system columns
	ORDER BY
		schema_name,
		table_name,
		col.attnum;`

	query := strings.Replace(queryTemplate, "[SCHEMAS]", getFormattedSchemaList(dbConnector.Input.Schemas), -1)
	return query
}

func (dbConnector PostgresDBConnector) GetRelationsQueryStatement() string {
	queryTemplate :=
		`SELECT
		tc.table_schema, 
		tc.constraint_name, 
		tc.table_name, 
		kcu.column_name, 
		ccu.table_schema AS foreign_table_schema,
		ccu.table_name AS foreign_table_name,
		ccu.column_name AS foreign_column_name 
	FROM 
		information_schema.table_constraints AS tc 
		JOIN information_schema.key_column_usage AS kcu
		  ON tc.constraint_name = kcu.constraint_name
		  AND tc.table_schema = kcu.table_schema
		JOIN information_schema.constraint_column_usage AS ccu
		  ON ccu.constraint_name = tc.constraint_name
		  AND ccu.table_schema = tc.table_schema
	WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_schema in ([SCHEMAS]);`

	query := strings.Replace(queryTemplate, "[SCHEMAS]", getFormattedSchemaList(dbConnector.Input.Schemas), -1)
	return query
}
