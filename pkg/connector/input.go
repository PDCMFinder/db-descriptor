package connector

type Input struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Schemas  []string
	Db       string
}
