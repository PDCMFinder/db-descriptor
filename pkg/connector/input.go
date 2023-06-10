package connector

/*
Input parameters.

A struct grouping the parameters for the main program, like the db credentials and the schemas names to process.
*/
type Input struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Schemas  []string
	Db       string
}
