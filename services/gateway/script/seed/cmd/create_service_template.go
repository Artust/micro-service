package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) CreateServiceTemplate(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (s:ServiceTemplate) WHERE s.name = $name AND s.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Default service template",
	})
	if err != nil {
		log.Error("error when match service template, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record service template, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`CREATE (s:ServiceTemplate{name: $name, detail:$detail, type:$type, corporationId:$corporationId, defaultRoutineId:$defaultRoutineId, 
			defaultAvatarId:$defaultAvatarId, routineIds:$routineIds, avatarIds:$avatarIds, createdBy:$createdBy, updatedBy:$updatedBy, 
			createdAt:datetime($createdAt), deletedAt:$deletedAt, updatedAt:$updatedAt})`, map[string]interface{}{
			"name":             "Default service template",
			"detail":           "DetailDefault",
			"type":             "warranty",
			"corporationId":    1,
			"defaultRoutineId": 2,
			"defaultAvatarId":  3,
			"routineIds":       []int64{4, 5, 6, 7},
			"avatarIds":        []int64{4, 2, 6, 8},
			"createdBy":        6,
			"updatedBy":        6,
			"createdAt":        time.Now().Format(time.RFC3339),
			"updatedAt":        "",
			"deletedAt":        "",
		})
		if err != nil {
			log.Error("error when match routine, error: ", err)
			return err
		}
	}
	return nil
}
