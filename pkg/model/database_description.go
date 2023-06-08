package model

import (
	"strings"
)

type DatabaseDescription struct {
	Schemas []Schema
}

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
