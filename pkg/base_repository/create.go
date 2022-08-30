package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) Create(neo4jTransaction neo4j.Transaction, data interface{}) error {
	input := make(map[string]interface{})
	//Check type
	typePointer := reflect.ValueOf(data).Kind()
	if typePointer != reflect.Ptr {
		return errors.New("invalid pointer type")
	}
	labelName := getStructName(data)
	acronymName := strings.ToLower(labelName[:1])
	createQueryString := fmt.Sprintf(`CREATE (%s: %s{`, acronymName, labelName)
	returnQueryString := fmt.Sprintf(` RETURN id(%s), `, acronymName)
	value := reflect.ValueOf(data).Elem()
	for i := 0; i < value.NumField(); i++ {
		// Get the field tag value
		tag := value.Type().Field(i).Tag.Get("neo4j")
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		if tag == string(IdTag) ||
			tag == string(CreatedAtTag) ||
			tag == string(UpdatedAtTag) ||
			tag == string(DeletedAtTag) {
			continue
		}
		timeValue, ok := value.Field(i).Interface().(time.Time)
		if ok {
			createQueryString += fmt.Sprintf(`%s: datetime($%s),`, tag, tag)
			returnQueryString += fmt.Sprintf(`%s.%s, `, acronymName, tag)
			input[tag] = timeValue.Format(time.RFC3339)
			continue
		}
		createQueryString += fmt.Sprintf(`%s: $%s,`, tag, tag)
		returnQueryString += fmt.Sprintf(`%s.%s, `, acronymName, tag)
		input[tag] = value.Field(i).Interface()
	}
	createQueryString += `createdAt: datetime($createdAt), updatedAt:datetime($updatedAt), deletedAt:""})`
	returnQueryString += fmt.Sprintf(`%s.createdAt, `, acronymName)
	returnQueryString += fmt.Sprintf(`%s.updatedAt`, acronymName)
	input["createdAt"] = time.Now().Format(time.RFC3339)
	input["updatedAt"] = time.Now().Format(time.RFC3339)
	queryString := createQueryString + returnQueryString
	fmt.Println(queryString)
	fmt.Println(input)
	result, err := neo4jTransaction.Run(queryString, input)
	if err != nil {
		fmt.Printf("error when create %s, error: %v", labelName, err)
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
	return nil
}
