package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreatePermission(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (p:Permission) WHERE p.entity= $entity AND p.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"entity": "admin",
	})
	if err != nil {
		log.Error("error when match permission, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record permission, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (c:Permission{entity:$entity, permissionAction:$permissionAction,
						createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"entity":           "admin",
			"permissionAction": "Crud pos",
			"createdAt":        time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create permission, error: ", err)
			return err
		}
	}
	return nil
}
