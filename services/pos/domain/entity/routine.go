package entity

import "time"

type PosRoutine struct {
	Id                       int64     `neo4j:"id"`
	Name                     string    `neo4j:"name"`
	Detail                   string    `neo4j:"detail"`
	AnimationFile            string    `neo4j:"animationFile"`
	ImageFile                string    `neo4j:"imageFile"`
	SoundFile                string    `neo4j:"soundFile"`
	StartDate                time.Time `neo4j:"startDate"`
	EndDate                  time.Time `neo4j:"endDate"`
	PosId                    int64     `neo4j:"posId"`
	CategoryId               int64     `neo4j:"categoryId"`
	ServiceTemplateRoutineId int64     `neo4j:"serviceTemplateRoutineId"`
	Gender                   int64     `neo4j:"gender"`
	CreatedAt                time.Time `neo4j:"createdAt"`
	UpdatedAt                time.Time `neo4j:"updatedAt"`
	DeletedAt                time.Time `neo4j:"deletedAt"`
}

type ListPosRoutineByCategory struct {
	Results []*Routines `neo4j:"result"`
}

type Routines struct {
	Category    []*RoutineCategory `neo4j:"category"`
	ListRoutine []*PosRoutine      `neo4j:"listRoutine"`
}

type GetListRoutineOption struct {
	Page    int64
	PerPage int64
	PosId   int64
	Gender  int64
	Ids     []int64
}

type GetListRoutineByCategoryOption struct {
	Page           int64
	PerPage        int64
	PosId          int64
	ListCategoryId []int64
	Ids            []int64
	EndDate        time.Time
	Between        int64
}

type TrigerEventPayload struct {
	ID            int64
	AnimationFile string
	SoundFile     string
}

type TriggerEventOperatorSidePayload struct {
	PosID     int64
	RoutineID int64
}

type DeleteRoutineOption struct {
	Id    int64
	PosId int64
}
