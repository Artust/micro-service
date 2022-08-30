package entity

import "time"

type Note struct {
	Id            int64     `neo4j:"id"`
	TalkSessionId int64     `neo4j:"talkSessionId"`
	Content       string    `neo4j:"content"`
	IsGuest       bool      `neo4j:"isGuest"`
	CreatedAt     time.Time `neo4j:"createdAt"`
	UpdatedAt     time.Time `neo4j:"updatedAt"`
	DeletedAt     time.Time `neo4j:"deletedAt"`
}

type GetListNoteOption struct {
	Page          int64
	PerPage       int64
	TalkSessionId int64
}
