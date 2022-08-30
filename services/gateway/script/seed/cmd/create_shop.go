package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) CreateShop(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (s:Shop) WHERE s.name= $name AND s.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "Aeon",
	})
	if err != nil {
		log.Error("error when match shop, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record create shop, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (s:Shop{name:$name, address:$address,
			createBy:$createBy, createdAt:$createdAt, updatedAt:"", deletedAt:""}) RETURN id(s), s.name, s.address, 
			 s.createBy, s.createdAt, s.updatedAt`, map[string]interface{}{
			"name":      "Aeon",
			"address":   "Ha Noi, Viet Nam",
			"createBy":  1,
			"createdAt": time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create shop, error: ", err)
			return err
		}
	}
	return nil
}
