package entity

import "time"

type UserActivity struct {
	Id           int64     `neo4j:"id"`
	Description  string    `neo4j:"description"`
	ActivityName string    `neo4j:"activityName"`
	AccountId    int64     `neo4j:"accountId"`
	CreatedAt    time.Time `neo4j:"createdAt"`
	UpdatedAt    time.Time `neo4j:"updatedAt"`
	DeletedAt    time.Time `neo4j:"deletedAt"`
}
