package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) CreateTalkSession(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (t:TaklSession) WHERE t.storageLink= $storageLink AND t.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"storageLink": "",
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
		_, err := tx.Run(`create (t:TaklSession{storageLink:$storageLink, startTime:$startTime,endTime:$endTime,usedHotKey:$usedHotKey, sessionStatus:$sessionStatus,
			customerRecord:$customerRecord, avatarId:$avatarId, ipCameraId:$ipCameraId,conversation:$conversation, posId:$posId,
				createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"storageLink":      "",
			"startTime":        time.Now().Format(time.RFC3339),
			"endTime":          time.Now().Format(time.RFC3339),
			"usedHotKey":       "usedHotKey",
			"sessionStatus":    0,
			"customerRecord":   "",
			"avatarId":         112,
			"ipCameraId":       10,
			"conversation":     "Operator: hello how are you",
			"posId":            1,
			"permissionAction": "Crud pos",
			"createdAt":        time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create taklSession, error: ", err)
			return err
		}
	}
	return nil
}
