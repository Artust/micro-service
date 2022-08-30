package neo4j

import (
	"avatar/services/gateway/script/seed/config"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func ConnectNeo4j(environment *config.Environment) neo4j.Driver {
	driver, err := neo4j.NewDriver(environment.Neo4JConf.NEO4J_URI, neo4j.BasicAuth(environment.Neo4JConf.UserName, environment.Neo4JConf.Password, ""))
	if err != nil {
		panic(err)
	}
	return driver
}
