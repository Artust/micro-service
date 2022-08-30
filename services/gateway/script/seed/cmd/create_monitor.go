package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateMonitor(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (c:CustomerMonitor) WHERE c.maker= $maker AND c.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"maker": "sony",
	})
	if err != nil {
		log.Error("error when match customer monitor, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record create customer monitor, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (c:CustomerMonitor{maker:$maker, serialNumber:$serialNumber,
			resolutionWidth:$resolutionWidth, resolutionHeight:$resolutionHeight,rotation:$rotation,monitorStatus:$monitorStatus, avatarId:$avatarId,
			 createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"maker":            "sony",
			"serialNumber":     "Detail customer monitor",
			"resolutionWidth":  "1920",
			"resolutionHeight": "1080",
			"rotation":         0,
			"monitorStatus":    "standby",
			"avatarId":         1,
			"createdAt":        time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create customer monitor, error: ", err)
			return err
		}
	}
	return nil
}
