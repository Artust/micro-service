package entity

import "time"

type Monitor struct {
	Id                 int64     `neo4j:"id"`
	Maker              string    `neo4j:"maker"`
	SerialNumber       string    `neo4j:"serialNumber"`
	MonitorStatus      string    `neo4j:"monitorStatus"`
	ResolutionWidth    int64     `neo4j:"resolutionWidth"`
	ResolutionHeight   int64     `neo4j:"resolutionHeight"`
	HorizontalRotation bool      `neo4j:"horizontalRotation"`
	PosId              int64     `neo4j:"posId"`
	CreatedAt          time.Time `neo4j:"createdAt"`
	UpdatedAt          time.Time `neo4j:"updatedAt"`
	DeletedAt          time.Time `neo4j:"deletedAt"`
}

type GetListMonitorOption struct {
	Page     int64
	PerPage  int64
	PosId    int64
}
