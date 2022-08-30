package entity

import "time"

type Device struct {
	Id           int64     `neo4j:"id"`
	Maker        string    `neo4j:"maker"`
	SerialNumber string    `neo4j:"serialNumber"`
	DeviceType   string    `neo4j:"deviceType"`
	UsePurpose   string    `neo4j:"usePurpose"`
	Owner        int64     `neo4j:"owner"`
	OnsiteType   bool      `neo4j:"onsiteType"`
	AccountId    int64     `neo4j:"accountId"`
	PosId        int64     `neo4j:"posId"`
	Resolution   string    `neo4j:"resolution"`
	CenterId     int64     `neo4j:"centerId"`
	CreatedAt    time.Time `neo4j:"createdAt"`
	UpdatedAt    time.Time `neo4j:"updatedAt"`
	DeletedAt    time.Time `neo4j:"deletedAt"`
}

type GetListDeviceOption struct {
	Page       int64
	PerPage    int64
	AccountId  int64
	PosId      int64
	CenterId   int64
	DeviceType string
}
