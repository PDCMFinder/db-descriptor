// Package report contains logic to report the descriptions of a database.
// For now, a service to write a JSON file with the database descritions is provided.
package report

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/PDCMFinder/db-descriptor/pkg/model"
)

// Writes a [model/DatabaseDescription] struct as a JSOn file .
func WriteDbDescriptionAsJson(databaseDescription model.DatabaseDescription, outputFileName string) {
	jsonData, err := json.MarshalIndent(databaseDescription, "", "    ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Open a file for writing
	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON to file:", err)
	}

	fmt.Println("JSON file created successfully.")
}
