package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) FindOne(neo4jTransaction neo4j.Transaction, destination interface{}) error {
	if br.BaseModel == nil {
		return errors.New("missing model")
	}
	if len(br.Condition) == 0 {
		return errors.New("missing condition")
	}
	//Check type
	typePointer := reflect.ValueOf(destination).Kind()
	if typePointer != reflect.Ptr {
		return errors.New("invalid pointer type")
	}
	br.SkipParam = 0
	br.LimitParam = 1
	labelName := getStructName(br.BaseModel)
	acronymName := strings.ToLower(labelName[:1])
	matchQueryString := fmt.Sprintf(`MATCH (%s:%s) WHERE`, acronymName, labelName)
	conditionQuery := make(map[string]interface{}, len(br.Condition))
	for condition, value := range br.Condition {
		conditionSplit := strings.Split(condition, " ")
		if len(conditionSplit) < 3 {
			return errors.New("invalid condition")
		}
		field := conditionSplit[0]
		operator := conditionSplit[1]
		if field == "id" {
			matchQueryString += fmt.Sprintf(` id(%s) %s $id AND `, acronymName, operator)
		} else {
			matchQueryString += fmt.Sprintf(` %s.%s %s $%s AND `, acronymName, field, operator, field)
		}
		conditionQuery[field] = value
	}
	matchQueryString += fmt.Sprintf(`%s.deletedAt = ""`, acronymName)
	value := reflect.ValueOf(br.BaseModel).Elem()
	returnQueryString := fmt.Sprintf(` RETURN id(%s), `, acronymName)
	for i := 0; i < value.NumField(); i++ {
		// Get the field tag value
		tag := value.Type().Field(i).Tag.Get("neo4j")
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		if tag == string(IdTag) || tag == string(DeletedAtTag) {
			continue
		}
		if tag == string(UpdatedAtTag) {
			returnQueryString += fmt.Sprintf(`%s.%s`, acronymName, tag)
			continue
		}
		returnQueryString += fmt.Sprintf(`%s.%s, `, acronymName, tag)
	}
	if br.OrderParam != "" {
		returnQueryString += fmt.Sprintf(` ORDER BY %s.%s`, acronymName, br.OrderParam)
	} else {
		returnQueryString += fmt.Sprintf(` ORDER BY %s.createdAt DESC`, acronymName)
	}
	if br.SkipParam > 0 {
		returnQueryString += fmt.Sprintf(` SKIP %d`, br.SkipParam)
	}
	if br.LimitParam > 0 {
		returnQueryString += fmt.Sprintf(` LIMIT %d`, br.LimitParam)
	}
	queryString := matchQueryString + returnQueryString
	fmt.Println(queryString)
	fmt.Println(conditionQuery)
	result, err := neo4jTransaction.Run(queryString, conditionQuery)
	if err != nil {
		fmt.Printf("error when get %s, error: %v", labelName, err)
		return err
	}
	record, err := result.Single()
	if err != nil {
		fmt.Printf("error when record %s, error: %v", labelName, err)
		return err
	}
	for i := 0; i < value.NumField(); i++ {
		// Get the field tag value
		tag := value.Type().Field(i).Tag.Get("neo4j")
		if tag == string(IdTag) {
			createdValue, ok := record.Get(fmt.Sprintf("id(%s)", acronymName))
			if ok {
				value.Field(i).Set(reflect.ValueOf(createdValue))
			}
			continue
		}
		if tag == string(DeletedAtTag) {
			continue
		}
		createdValue, ok := record.Get(fmt.Sprintf("%s.%s", acronymName, tag))
		if ok {
			if value.Field(i).Kind() == reflect.Slice {
				slice := reflect.MakeSlice(value.Field(i).Type(), 0, 0)
				for _, val := range createdValue.([]interface{}) {
					slice = reflect.Append(slice, reflect.ValueOf(val))
				}
				value.Field(i).Set(slice)
			} else {
				value.Field(i).Set(reflect.ValueOf(createdValue))
			}
		}
	}
	destinationValue := reflect.ValueOf(destination)
	destinationValue.Elem().Set(value)
	return nil
}
