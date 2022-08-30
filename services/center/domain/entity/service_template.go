package entity

import "time"

type ServiceTemplate struct {
	Id                          int64     `neo4j:"id"`
	Name                        string    `neo4j:"name"`
	Detail                      string    `neo4j:"detail"`
	Type                        string    `neo4j:"type"`
	CorporationId               int64     `neo4j:"corporationId"`
	DefaultMaleRoutineIds       []int64   `neo4j:"DefaultMaleRoutineIds"`
	DefaultMaleRoutineListIds   []int64   `neo4j:"defaultMaleRoutineListIds"`
	MaleRoutineIds              []int64   `neo4j:"maleRoutineIds"`
	DefaultFemaleRoutineIds     []int64   `neo4j:"defaultFemaleRoutineIds"`
	DefaultFemaleRoutineListIds []int64   `neo4j:"defaultFemaleRoutineListIds"`
	FemaleRoutineIds            []int64   `neo4j:"femaleRoutineIds"`
	DefaultMaleAvatarId         int64     `neo4j:"defaultMaleAvatarId"`
	DefaultFemaleAvatarId       int64     `neo4j:"defaultFemaleAvatarId"`
	AvatarIds                   []int64   `neo4j:"avatarIds"`
	UpdatedBy                   int64     `neo4j:"updatedBy"`
	CreatedBy                   int64     `neo4j:"createdBy"`
	ServiceTemplateCategory     int64     `neo4j:"serviceTemplateCategory"`
	Backgrounds                 []string  `neo4j:"backgrounds"`
	BackgroundDefault           string    `neo4j:"backgroundDefault"`
	CreatedAt                   time.Time `neo4j:"createdAt"`
	UpdatedAt                   time.Time `neo4j:"updatedAt"`
	DeletedAt                   time.Time `neo4j:"deletedAt"`
}

type GetListServiceTemplateOption struct {
	Page                    int64
	PerPage                 int64
	CorporationId           int64
	ServiceTemplateCategory int64
}
