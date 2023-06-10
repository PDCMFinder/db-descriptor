// Package model contains structs to represent the information extracted from a database, like
// table names, view names, column names, etc.
package model

import (
	"strings"
)

/*
A container for the different Schemas for which descriptions where extracted.

DatabaseDescription contains a slice of `Schema`.
*/
type DatabaseDescription struct {
	Schemas []Schema
}

// Returns a string representation of the DatabaseDescription struct.
func (d DatabaseDescription) String() string {
	var sb strings.Builder

	for _, schema := range d.Schemas {
		sb.WriteString(schema.Name + "\n")
		for _, entity := range schema.Entities {
			sb.WriteString("\t" + entity.Name + "\n")
			for _, column := range entity.Columns {
				sb.WriteString("\t\t" + column.Name + "\n")
			}
		}
	}

	return sb.String()
}
