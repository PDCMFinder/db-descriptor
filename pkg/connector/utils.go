package connector

import "strings"

func GetFormattedSchemaList(schemaList []string) string {
	formattedSchemasList := make([]string, 0)
	for _, e := range schemaList {
		quoted := "'" + e + "'"
		formattedSchemasList = append(formattedSchemasList, quoted)
	}
	return strings.Join(formattedSchemasList, ",")
}
