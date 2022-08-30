package main

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c Server) CreateIpCamera(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (i:IpCamera) WHERE i.name= $maker AND i.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"maker": "panasonic",
	})
	if err != nil {
		log.Error("error when match ipcamera, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record ipcamera, error: ", err)
		return err
	}
	if record.Values[0].(int64) == 0 {
		_, err := tx.Run(`create (c:IpCamera{maker:$maker, serialNumber:$serialNumber,publicIpCamera:$publicIpCamera, privateIpCamera:$privateIpCamera,
		resolutionWidth:$resolutionWidth, resolutionHeight:$resolutionHeight,cameraStatus:$cameraStatus, customerMonitorId:$customerMonitorId,
		createdAt:$createdAt, updatedAt:"", deletedAt:""})`, map[string]interface{}{
			"maker":             "panasonic",
			"serialNumber":      "Detail customer monitor",
			"publicIpCamera":    "192.168.1.1",
			"privateIpCamera":   "192.168.2.1",
			"resolutionWidth":   "1920",
			"resolutionHeight":  "1080",
			"cameraStatus":      "active",
			"customerMonitorId": 1,
			"createdAt":         time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error create ipcamera, error: ", err)
			return err
		}
	}
	return nil
}
