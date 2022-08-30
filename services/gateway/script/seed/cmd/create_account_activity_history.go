package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateAccountActivityHistory(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (a:AccountActivityHistory) WHERE a.name= $name AND a.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Dealing with customers",
	})
	if err != nil {
		log.Error("error when match account activity history, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record account activity history, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (t:AccountActivityHistory{name:$name, accountId:$accountId,
			createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"name":      "Dealing with customers",
			"accountId": 25,
			"createdAt": time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create account activity history, error: ", err)
			return err
		}
	}
	return nil
}
