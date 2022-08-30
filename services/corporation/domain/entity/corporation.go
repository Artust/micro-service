package entity

import "time"

type Corporation struct {
	Id        int64     `neo4j:"id"`
	Name      string    `neo4j:"name"`
	Address   string    `neo4j:"address"`
	Detail    string    `neo4j:"detail"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type GetListCorporationOption struct {
	Page      int64
	PerPage   int64
}
