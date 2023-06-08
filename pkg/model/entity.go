package model

import "fmt"

type Entity struct {
	SchemaName string
	Name       string
	EntityType string
	Columns    []Column
	Comment    string
}

func (e Entity) String() string {
	return fmt.Sprintf("[%+v, %s, %s, %s, %s]", e.SchemaName, e.Name, e.EntityType, e.Columns, e.Comment)
}
