package entity

import "time"

type CenterRoutine struct {
	Id                int64     `neo4j:"id"`
	Name              string    `neo4j:"name"`
	Detail            string    `neo4j:"detail"`
	AnimationFile     string    `neo4j:"animationFile"`
	ImageFile         string    `neo4j:"imageFile"`
	SoundFile         string    `neo4j:"soundFile"`
	StartDate         time.Time `neo4j:"startDate"`
	EndDate           time.Time `neo4j:"endDate"`
	CategoryId        int64     `neo4j:"categoryId"`
	Gender            int64     `neo4j:"gender"`
	CreatedAt         time.Time `neo4j:"createdAt"`
	UpdatedAt         time.Time `neo4j:"updatedAt"`
	DeletedAt         time.Time `neo4j:"deletedAt"`
}

type GetListRoutineOption struct {
	Page              int64
	PerPage           int64
	CategoryId        int64
	Gender            int64
	CategoryIds       []int64
	Ids               []int64
}
