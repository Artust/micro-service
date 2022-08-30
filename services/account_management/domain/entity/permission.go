package entity

import "time"

type Permission struct {
	Id               int64     `neo4j:"id"`
	Entity           string    `neo4j:"entity"`
	PermissionAction string    `neo4j:"permissionAction"`
	CreatedAt        time.Time `neo4j:"createdAt"`
	UpdatedAt        time.Time `neo4j:"updatedAt"`
	DeletedAt        time.Time `neo4j:"deletedAt"`
}
