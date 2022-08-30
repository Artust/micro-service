package main

import (
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateAvatar(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (a:Avatar) WHERE a.name= $name AND a.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Avatar 1",
	})
	if err != nil {
		log.Error("error when match avatar, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record create avatar, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (s:Avatar{name:$name, detail:$detail,
			imageLink:$imageLink, vrmLink:$vrmLink,startDate:datetime($startDate),endDate:datetime($endDate), createdAt:$createdAt, 
			updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"name":      "Avatar 1",
			"detail":    "Detail avatar 1",
			"imageLink": fmt.Sprintf("/%v/%v", "routine", c.urlImageDefault),
			"vrmLink":   "",
			"startDate": time.Now().Format(time.RFC3339),
			"endDate":   time.Now().Format(time.RFC3339),
			"createdAt": time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create avatar, error: ", err)
			return err
		}
	}
	return nil
}
