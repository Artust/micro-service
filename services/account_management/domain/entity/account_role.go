package entity

import "time"

type AccountRole struct {
	Id            int64     `neo4j:"id"`
	Name          string    `neo4j:"name"`
	PermissionIds []int64   `neo4j:"permissionIds"`
	Level         int64     `neo4j:"level"`
	CreatedAt     time.Time `neo4j:"createdAt"`
	UpdatedAt     time.Time `neo4j:"updatedAt"`
	DeletedAt     time.Time `neo4j:"deletedAt"`
}
