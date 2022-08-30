package neo4j

import (
	"avatar/services/pos/config"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func ConnectNeo4j(environment *config.Environment) neo4j.Driver {
	driver, err := neo4j.NewDriver(environment.Neo4jUri,
		neo4j.BasicAuth(environment.Neo4jUserName, environment.Neo4jPassword, ""))
	if err != nil {
		panic(err)
	}
	return driver
}
