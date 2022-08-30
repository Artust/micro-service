package main

import (
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) CreateRoutine(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (r:Routine) WHERE r.name = $name AND r.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "NameDefault",
	})
	if err != nil {
		log.Error("error when match routine, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record routine category, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run("CREATE (r:Routine{name: $name, detail:$detail, startDate:datetime($startDate), endDate:datetime($endDate), animationKey:$animationKey, soundFile:$soundFile, createdAt:datetime($createdAt), deletedAt:$deletedAt, updatedAt:$updatedAt, image: $image}) WITH r MATCH (c:RoutineCategory) WHERE id(c) = $categoryId CREATE (r)-[b:BELONG_ROUTINE_CATEGORY]->(c) return r", map[string]interface{}{
			"name":         "NameDefault",
			"detail":       "DetailDefault",
			"animationKey": fmt.Sprintf("/%v/%v", "routine", c.urlAnimationKey),
			"soundFile":    fmt.Sprintf("/%v/%v", "routine", c.urlSound),
			"image":        fmt.Sprintf("/%v/%v", "routine", c.urlImageDefault),
			"startDate":    time.Now().Format(time.RFC3339),
			"endDate":      time.Now().Format(time.RFC3339),
			"createdAt":    time.Now().Format(time.RFC3339),
			"updatedAt":    "",
			"deletedAt":    "",
			"categoryId":   "CategoryIdDefault",
		})
		if err != nil {
			log.Error("error when match routine, error: ", err)
			return err
		}
	}
	return nil
}
