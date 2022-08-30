package entity

import "time"

type Center struct {
	Id            int64     `neo4j:"id"`
	Name          string    `neo4j:"name"`
	Type          string    `neo4j:"type"`
	Detail        string    `neo4j:"detail"`
	CorporationId int64     `neo4j:"corporationId"`
	CreatedAt     time.Time `neo4j:"createdAt"`
	UpdatedAt     time.Time `neo4j:"updatedAt"`
	DeletedAt     time.Time `neo4j:"deletedAt"`
}

type GetListCenterOption struct {
	Page          int64
	PerPage       int64
	CorporationId int64
}
