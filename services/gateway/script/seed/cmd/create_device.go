package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateDevice(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (a:Device) WHERE a.maker= $maker AND a.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"maker": "sony",
	})
	if err != nil {
		log.Error("error when match taklSession, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record taklSession, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (t:Device{maker:$maker, serialNumber:$serialNumber,deviceType:$deviceType, usePurpose : $usePurpose,
			owner:$owner, user:$user, onsiteType:$onsiteType, accountId:$accountId, POSId:$POSId, centerId:$centerId,
					createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"maker":        "sony",
			"serialNumber": "SACB11246",
			"deviceType":   "camera",
			"usePurpose":   "usePurpose",
			"owner":        "owner",
			"user":         "user",
			"onsiteType":   "onsiteType",
			"POSId":        15,
			"accountId":    25,
			"centerId":     10,
			"createdAt":    time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create device, error: ", err)
			return err
		}
	}
	return nil
}
