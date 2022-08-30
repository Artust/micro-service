package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateCenter(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (c:Center) WHERE c.name= $name AND c.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Aeon",
	})
	if err != nil {
		log.Error("error when match center, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record create center, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (s:Center{name:$name, detail:$detail,
			type:$type, createdAt:$createdAt, updatedAt:"", deletedAt:""}) RETURN id(s), s.name, s.address, 
			 s.createBy, s.createdAt, s.updatedAt`, map[string]interface{}{
			"name":      "Aeon",
			"detail":    "shopping mall",
			"type":      1,
			"createdAt": time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create center, error: ", err)
			return err
		}
	}
	return nil
}
