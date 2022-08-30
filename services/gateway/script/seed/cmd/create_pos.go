package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreatePos(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (p:POS) WHERE p.serviceName= $serviceName AND p.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"serviceName": "service name default",
	})
	if err != nil {
		log.Error("error when match pos, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record pos, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err = tx.Run(`CREATE (p:POS{serviceName:$serviceName, serviceType:$serviceType, serviceDetail:$serviceDetail, shopId:$shopId, 
		centerId:$centerId, serviceTemplateId:$serviceTemplateId, defaultRoutineId:$defaultRoutineId, defaultAvatarId:$defaultAvatarId, routineIds:$routineIds, 
		avatarIds:$avatarIds, createdBy:$createdBy, createdAt:datetime($createdAt), updatedAt:$updatedAt, deletedAt:$deletedAt, 
		standbyRoutineId:$standbyRoutineId,greetingRoutineId:$greetingRoutineId})`, map[string]interface{}{
			"serviceName":       "service name default",
			"serviceType":       "service type default",
			"serviceDetail":     "service detail defaults",
			"shopId":            1,
			"centerId":          2,
			"serviceTemplateId": 10,
			"defaultRoutineId":  11,
			"defaultAvatarId":   12,
			"routineIds":        []int64{1, 2, 3, 4, 5},
			"avatarIds":         []int64{1, 2, 3, 4, 5},
			"createdBy":         1,
			"standbyRoutineId":  10,
			"greetingRoutineId": 15,
			"updatedAt":         "",
			"createdAt":         time.Now().Format(time.RFC3339),
			"deletedAt":         "",
		})
		if err != nil {
			log.Error("error create pos, error: ", err)
			return err
		}
	}
	return nil
}
