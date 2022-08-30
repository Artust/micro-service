package base_repository

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type CreateRelationshipData struct {
	SrcId            int64
	SrcLabel         string
	RelationshipName string
	DstId            int64
	DstLabel         string
}

func (br *BaseRepository) CreateRelationship(neo4jTransaction neo4j.Transaction, data CreateRelationshipData) error {
	createQueryString := fmt.Sprintf(
		`MATCH (src:%s),(dst:%s) WHERE id(src) = %d AND id(dst) = %d CREATE (src)-[r:%s]->(dst)`,
		data.SrcLabel,
		data.DstLabel,
		data.SrcId,
		data.DstId,
		data.RelationshipName,
	)
	fmt.Println(createQueryString)
	_, err := neo4jTransaction.Run(createQueryString, map[string]interface{}{})
	if err != nil {
		fmt.Printf("error when create relationship, error: %v", err)
		return err
	}
	return nil
}

func (br *BaseRepository) DeleteRelationship(neo4jTransaction neo4j.Transaction, srcId int64, dstId int64) error {
	deleteQueryString := fmt.Sprintf(
		`MATCH (src)-[r]-(dst) WHERE id(src) = %d AND id(dst) = %d DELETE r`,
		srcId,
		dstId,
	)
	fmt.Println(deleteQueryString)
	_, err := neo4jTransaction.Run(deleteQueryString, map[string]interface{}{})
	if err != nil {
		fmt.Printf("error when delete relationship, error: %v", err)
		return err
	}
	return nil
}
