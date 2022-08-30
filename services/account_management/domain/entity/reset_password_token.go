package entity

import "time"

type ResetPasswordToken struct {
	Id        int64     `neo4j:"id"`
	AccountId int64     `neo4j:"accountId"`
	Email     string    `neo4j:"email"`
	Username  string    `neo4j:"username"`
	Token     string    `neo4j:"token"`
	ExpiresAt time.Time `neo4j:"expiresAt"`
	CreatedAt time.Time `neo4j:"createdAt"`
	UpdatedAt time.Time `neo4j:"updatedAt"`
	DeletedAt time.Time `neo4j:"deletedAt"`
}

type ResetPassword struct {
	Username    string
	ResetPwdURL string
	Token       string
}
