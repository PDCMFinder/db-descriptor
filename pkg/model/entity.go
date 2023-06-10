package model

import "fmt"

/*
A representation of a database entity (table, view, for example).

Entity struct contains data that can be extracted from the database, like the name and the comment. It also has a slice
of [Column].
*/
type Entity struct {
	SchemaName string
	Name       string
	EntityType string
	Columns    []Column
	Comment    string
}

// Returns a string representation of the Entity struct.
func (e Entity) String() string {
	return fmt.Sprintf("[%+v, %s, %s, %s, %s]", e.SchemaName, e.Name, e.EntityType, e.Columns, e.Comment)
}
