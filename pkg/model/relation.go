package model

/*
A representation of a relation between 2 entities.

The Relation struct represents data about a foreign key.
*/
type Relation struct {
	SchemaName          string
	EntityName          string
	RelationName        string
	ColumnName          string
	ForeignEntitySchema string
	ForeignEntityName   string
	ForeignColumnName   string
}
