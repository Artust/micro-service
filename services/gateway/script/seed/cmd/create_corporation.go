package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateCorporation(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (c:Corporation) WHERE c.name= $name AND c.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Default corporation",
	})
	if err != nil {
		log.Error("error when match corporation, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record corporation, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (c:Corporation{name:$name, address:$address,detail:$detail,
				createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"name":      "Default corporation",
			"address":   "Ha Noi, Viet Nam",
			"detail":    "Detail corporation",
			"createdAt": time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create corporation, error: ", err)
			return err
		}
	}
	return nil
}
