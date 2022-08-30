package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) Delete(neo4jTransaction neo4j.Transaction) (rowsAffected int64, err error) {
	if br.BaseModel == nil {
		return 0, errors.New("missing model")
	}
	if len(br.Condition) == 0 {
		return 0, errors.New("missing condition")
	}
	input := make(map[string]interface{})
	//Check type
	typePointer := reflect.ValueOf(br.BaseModel).Kind()
	if typePointer != reflect.Ptr {
		return 0, errors.New("invalid pointer type")
	}
	labelName := getStructName(br.BaseModel)
	acronymName := strings.ToLower(labelName[:1])
	matchQueryString := fmt.Sprintf(`MATCH (%s:%s) WHERE`, acronymName, labelName)
	for condition, value := range br.Condition {
		conditionSplit := strings.Split(condition, " ")
		if len(conditionSplit) < 3 {
			return 0, errors.New("invalid condition")
		}
		field := conditionSplit[0]
		operator := conditionSplit[1]
		if field == "id" {
			matchQueryString += fmt.Sprintf(` id(%s) %s $id AND `, acronymName, operator)
		} else {
			matchQueryString += fmt.Sprintf(` %s.%s %s $%s AND `, acronymName, field, operator, field)
		}
		input[field] = value
	}
	matchQueryString += fmt.Sprintf(`%s.deletedAt = ""`, acronymName)
	setString := fmt.Sprintf(` SET %s.deletedAt= datetime($deletedAt)`, acronymName)
	input["deletedAt"] = time.Now().Format(time.RFC3339)
	queryString := matchQueryString + setString
	fmt.Println(queryString)
	fmt.Println(input)
	result, err := neo4jTransaction.Run(queryString, input)
	if err != nil {
		fmt.Printf("error when delete %s, error: %v", labelName, err)
		return 0, err
	}
	record, err := result.Consume()
	if err != nil {
		fmt.Printf("error when delete %s, error: %v", labelName, err)
		return 0, err
	}
	return int64(record.Counters().PropertiesSet()), nil
}
