package connector

import "strings"

// Helper method to format correctly a list of schema names to be used as a filter in a SQL query.
func getFormattedSchemaList(schemaList []string) string {
	formattedSchemasList := make([]string, 0)
	for _, e := range schemaList {
		quoted := "'" + e + "'"
		formattedSchemasList = append(formattedSchemasList, quoted)
	}
	return strings.Join(formattedSchemasList, ",")
}
