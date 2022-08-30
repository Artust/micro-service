package entity

import "time"

type ServiceTemplateCategory struct {
	Id        int64     `neo4j:"id"`
	Name      string    `neo4j:"name"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type GetListServiceTemplateCategoryOption struct {
	Page    int64
	PerPage int64
}
