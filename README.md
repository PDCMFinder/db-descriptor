# DB Descriptor
DB Descriptor — a light tool for describing a database.

This tool inspects a database (for now PostgreSQL) and retrieves basic information about some database objects: table/view names, column names, column data type and comments. It writes a JSON file with the information, but also exposes methods to process the information programatically.


## Usage
```
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
```
