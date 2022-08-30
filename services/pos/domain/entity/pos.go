package entity

import "time"

type Pos struct {
	Id                      int64     `neo4j:"id"`
	Name                    string    `neo4j:"name"`
	ServiceName             string    `neo4j:"serviceName"`
	ServiceCategory         string    `neo4j:"serviceCategory"`
	ServiceDetail           string    `neo4j:"serviceDetail"`
	ShopId                  int64     `neo4j:"shopId"`
	CenterId                int64     `neo4j:"centerId"`
	ServiceTemplateId       int64     `neo4j:"serviceTemplateId"`
	MaleRoutineIds          []int64   `neo4j:"maleRoutineIds"`
	DefaultMaleRoutineIds   []int64   `neo4j:"defaultMaleRoutineIds"`
	FemaleRoutineIds        []int64   `neo4j:"femaleRoutineIds"`
	DefaultFemaleRoutineIds []int64   `neo4j:"defaultFemaleRoutineIds"`
	DefaultAvatarId         int64     `neo4j:"defaultAvatarId"`
	CreatedBy               int64     `neo4j:"createdBy"`
	UseServiceTemplate      bool      `neo4j:"useServiceTemplate"`
	Backgrounds             []string  `neo4j:"background"`
	DefaultBackground       string    `neo4j:"defaultBackground"`
	CreatedAt               time.Time `neo4j:"createdAt"`
	UpdatedAt               time.Time `neo4j:"updatedAt"`
	DeletedAt               time.Time `neo4j:"deletedAt"`
}

type GetListPosOption struct {
	Page     int64
	PerPage  int64
	CenterId int64
}
