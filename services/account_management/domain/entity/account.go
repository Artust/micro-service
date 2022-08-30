package entity

import "time"

type Account struct {
	Id        int64     `neo4j:"id"`
	Email     string    `neo4j:"email"`
	Username  string    `neo4j:"username"`
	Password  string    `neo4j:"password"`
	Gender    int64     `neo4j:"gender"`
	RoleId    int64     `neo4j:"roleId"`
	CenterId  int64     `neo4j:"centerId"`
	Status    int64     `neo4j:"status"`
	CreatedBy int64     `neo4j:"createdBy"`
	Avatar    string    `neo4j:"avatar"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type ChangePasswordData struct {
	Id          int64
	Password    string
	NewPassword string
}

type GetListAccountOption struct {
	Page     int64
	PerPage  int64
	Gender   int64
	RoleId   int64
	CenterId int64
	Status   int64
}
