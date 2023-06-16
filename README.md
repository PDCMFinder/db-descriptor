# DB Descriptor

[![License](https://img.shields.io/github/license/PDCMFinder/db-descriptor)](https://github.com/PDCMFinder/db-descriptor/blob/main/LICENSE)

DB Descriptor is a lightweight tool for describing a database. It allows you to inspect a database (currently PostgreSQL) and retrieve basic information about its objects such as table/view names, column names, column data types, and comments. The tool writes the information to a JSON file and provides methods for programmatic processing.

## Features

- Inspect a PostgreSQL database and retrieve essential information about its objects
- Retrieve table/view names, column names, column data types, and comments
- Generate a JSON file containing the database description
- Programmatically process the retrieved information

## Installation

To install DB Descriptor, you need to have Go installed. Run the following command:

```bash
go get github.com/PDCMFinder/db-descriptor
```

## Usage
```
USAGE:
   db-descriptor [global options] command [command options] [arguments...]

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

## Contributing
Contributions are welcome! If you find any issues or have suggestions, please open an issue or submit a pull request.

## License
This project is licensed under the Apache License 2.0.
