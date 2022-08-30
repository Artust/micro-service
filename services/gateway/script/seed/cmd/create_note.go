package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateNote(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (tn:TalkSessionNote) WHERE tn.content= $content AND tn.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"content": "I want to buy this item",
	})
	if err != nil {
		log.Error("error when match talkSessionNote, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record create note, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err = tx.Run("CREATE (tn:TalkSessionNote {talkSessionId: $id, content: $content, private: $private, createdAt: datetime($createdAt), updatedAt: $updatedAt, deletedAt: $deletedAt}) RETURN tn", map[string]interface{}{
			"id":        1,
			"content":   "I want to buy this item",
			"private":   true,
			"createdAt": time.Now().Format(time.RFC3339),
			"updatedAt": "",
			"deletedAt": "",
		})
		if err != nil {
			log.Error("error when match talkSessionNote, error: ", err)
			return err
		}
	}
	return nil
}
