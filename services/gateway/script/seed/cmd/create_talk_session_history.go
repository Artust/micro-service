package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) CreateTalkSessionHistory(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (t:TaklSessionHistory) WHERE t.activeType= $activeType AND t.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"activeType": "",
	})
	if err != nil {
		log.Error("error when match tall session history, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record talk session history, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (t:TaklSessionHistory{activeType:$activeType, startTime:$startTime,endTime:$endTime,usedHotKey:$usedHotKey, accountId:$accountId,
			talkSessionId:$talkSessionId,
			createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"activeType":    "",
			"startTime":     time.Now().Format(time.RFC3339),
			"endTime":       time.Now().Format(time.RFC3339),
			"usedHotKey":    "usedHotKey",
			"accountId":     0,
			"talkSessionId": 10,
			"createdAt":     time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create talk session history, error: ", err)
			return err
		}
	}
	return nil
}
