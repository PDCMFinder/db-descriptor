package model

import (
	"strings"
)

/*
A container for entities descritpions in the database.

Schema struct contains the name of the schema (or namespace) and the slice of [Entity] that belong to it.
*/
type Schema struct {
	Name     string
	Entities []Entity
}

// Helper function to filter entities by type ("view", "table")
func (s Schema) GetEntitiesByType(entityType string) []Entity {
	var filtered []Entity
	for _, e := range s.Entities {
		if strings.ToLower(entityType) == strings.ToLower(e.EntityType) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
