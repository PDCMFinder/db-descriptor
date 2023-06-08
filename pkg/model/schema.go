package model

import (
	"strings"
)

type Schema struct {
	Name     string
	Entities []Entity
}

func (s Schema) GetEntitiesByType(entityType string) []Entity {
	var filtered []Entity
	for _, e := range s.Entities {
		if strings.ToLower(entityType) == strings.ToLower(e.EntityType) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
