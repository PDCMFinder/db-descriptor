/*
db-descriptor analyses a list of schemas in a database and creates a JSON file with information about its tables, views and columns.
USAGE:

	database descriptor [global options] command [command options] [arguments...]

COMMANDS:

	help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:

	--host value, -H value                                   database host (default: "localhost")
	--port value, -P value                                   database port (default: 8080)
	--user value, -u value                                   database user (default: "admin")
	--password value, -p value                               database password (default: "password")
	--name value, -n value                                   database name (default: "test")
	--schemas value, -s value [ --schemas value, -s value ]  comma separated list of schemas to describe (default: "public")
	--dbtype value, --dt value                               specify the database type (default: "postgres")
	--output value, -o value                                 JSON output file name the description of the database (default: "output.json")
	--help, -h                                               show help
*/
package main

import (
	"log"
	"os"

	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/report"
	"github.com/PDCMFinder/db-descriptor/pkg/service"
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
	databaseDescription := service.GetDbDescription(input)
	report.WriteDbDescriptionAsJson(databaseDescription, outputFileName)
	return nil
}
