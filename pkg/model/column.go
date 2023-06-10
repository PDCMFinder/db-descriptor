package model

/*
A representation of a database column.

Column contains data that can be extracted from the database. Main data is the name, type and comment. The rest is to
be able to identify the Entity it belongs to.
*/
type Column struct {
	SchemaName string
	EntityName string
	Name       string
	DataType   string
	Comment    string
}
