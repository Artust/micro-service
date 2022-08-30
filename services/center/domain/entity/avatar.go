package entity

import "time"

type Avatar struct {
	Id        int64     `neo4j:"id"`
	Name      string    `neo4j:"name"`
	Detail    string    `neo4j:"detail"`
	Image     string    `neo4j:"image"`
	Vrm       string    `neo4j:"vrm"`
	StartDate time.Time `neo4j:"startDate"`
	EndDate   time.Time `neo4j:"endDate"`
	Gender    int64     `neo4j:"gender"`
	Version   string    `neo4j:"version"`
	Exporter  string    `neo4j:"exporter"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type GetListAvatarOption struct {
	Page    int64
	PerPage int64
	Gender  int64
	Ids     []int64
}
