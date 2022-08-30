package entity

import "time"

type IpCamera struct {
	Id              int64     `neo4j:"id"`
	IsPrimaryCamera bool      `neo4j:"isPrimaryCamera"`
	DeviceId        int64     `neo4j:"deviceId"`
	PublicURI       string    `neo4j:"publicURI"`
	PrivateURI      string    `neo4j:"privateURI"`
	CreatedAt       time.Time `neo4j:"createdAt"`
	UpdatedAt       time.Time `neo4j:"updatedAt"`
	DeletedAt       time.Time `neo4j:"deletedAt"`
}

type GetListIpCameraOption struct {
	Page     int64
	PerPage  int64
	PosId    int64
	DeviceId []int64
}
