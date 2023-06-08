package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PDCMFinder/db-descriptor/internal/extractor"
	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/report"
	"github.com/urfave/cli/v2"
)

func main() {
	var host string
	var port int
	var user string
	var password string
	var name string
	var schemas cli.StringSlice
	var dbtype string
	var output string

	app := &cli.App{
		Name:  "database descriptor",
		Usage: "describes a database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"H"},
				Value:       "localhost",
				Usage:       "database host",
				Destination: &host,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"P"},
				Value:       8080,
				Usage:       "database port",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Value:       "admin",
				Usage:       "database user",
				Destination: &user,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "password",
				Usage:       "database password",
				Destination: &password,
			},
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "test",
				Usage:       "database name",
				Destination: &name,
			},
			&cli.StringSliceFlag{
				Name:        "schemas",
				Aliases:     []string{"s"},
				Value:       cli.NewStringSlice("public"),
				Usage:       "comma separated list of schemas to describe",
				Destination: &schemas,
			},
			&cli.StringFlag{
				Name:        "dbtype",
				Aliases:     []string{"dt"},
				Value:       "postgres",
				Usage:       "specify the database type",
				Destination: &dbtype,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "output.json",
				Usage:       "JSON output file name the description of the database",
				Destination: &output,
			},
		},
		Action: func(cCtx *cli.Context) error {
			input := connector.Input{
				Host:     host,
				Port:     port,
				User:     user,
				Password: password,
				Name:     name,
				Schemas:  schemas.Value(),
				Db:       dbtype,
			}
			return RunDBDescriptor(input, cCtx.String("output"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RunDBDescriptor(input connector.Input, outputFileName string) error {
	dbConnector, err := getDBConnector(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbConnector)
	dbDescriptionExtractor := extractor.New(dbConnector)
	databaseDescription := dbDescriptionExtractor.ExtractDescription()
	report.WriteDbDescriptionAsJson(databaseDescription, outputFileName)
	return nil
}

func getDBConnector(input connector.Input) (connector.DBConnector, error) {
	var dbConnector connector.DBConnector
	switch input.Db {
	case "postgres":
		dbConnector = connector.PostgresDBConnector{Input: input}
	default:
		return nil, errors.New(fmt.Sprintf("Database type [%s] not supported.", input.Db))
	}

	return dbConnector, nil
}
