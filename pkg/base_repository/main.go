package base_repository

import (
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type BaseRepository struct {
	BaseModel   interface{}
	Data        interface{}
	Condition   map[string]interface{}
	SkipParam   int64
	LimitParam  int64
	OrderParam  string
	SelectParam map[string]bool
	Neo4j       neo4j.Driver
}

func CreateBaseRepository(neo4j neo4j.Driver) *BaseRepository {
	return &BaseRepository{
		Neo4j: neo4j,
	}
}

func getStructName(data interface{}) string {
	valueOf := reflect.ValueOf(data)
	if valueOf.Type().Kind() == reflect.Ptr {
		if valueOf.Elem().Kind() == reflect.Slice {
			return valueOf.Elem().Type().Elem().Name()
		}
		return reflect.Indirect(valueOf).Type().Name()
	}
	return valueOf.Type().Name()
}

type StructTag string

var (
	CreatedAtTag StructTag = "createdAt"
	UpdatedAtTag StructTag = "updatedAt"
	DeletedAtTag StructTag = "deletedAt"
	IdTag        StructTag = "id"
)
