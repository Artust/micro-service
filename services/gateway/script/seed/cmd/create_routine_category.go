package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateRoutineCategory(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (rc:RoutineCategory) WHERE rc.name = $name AND rc.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Greeting",
	})
	if err != nil {
		log.Error("error when match routine category, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record routine category, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err = tx.Run("CREATE (rc:RoutineCategory{name: $name, createdAt:datetime($createdAt), updatedAt:$updatedAt, deletedAt:$deletedAt}) return id(rc), rc.name, rc.createdAt, rc.updatedAt, rc.deletedAt", map[string]interface{}{
			"name":      "Greeting",
			"createdAt": time.Now().Format(time.RFC3339),
			"updatedAt": "",
			"deletedAt": "",
		})
		if err != nil {
			log.Error("error when match routine, error: ", err)
			return err
		}
	}
	return nil
}
