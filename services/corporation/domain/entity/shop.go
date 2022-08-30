package entity

import "time"

type Shop struct {
	Id        int64     `neo4j:"id"`
	Name      string    `neo4j:"name"`
	Address   string    `neo4j:"address"`
	CreatedBy int64     `neo4j:"createdBy"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type GetListShopOption struct {
	Page    int64
	PerPage int64
}
