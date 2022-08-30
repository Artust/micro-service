package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) Select(fields ...string) *BaseRepository {
	br.SelectParam = make(map[string]bool)
	for _, field := range fields {
		if field == "*" {
			value := reflect.ValueOf(br.BaseModel).Elem()
			for i := 0; i < value.NumField(); i++ {
				// Get the field tag value
				tag := value.Type().Field(i).Tag.Get("neo4j")
				// Skip if tag is not defined or ignored
				if tag == "" || tag == "-" {
					continue
				}
				if tag == string(IdTag) ||
					tag == string(DeletedAtTag) ||
					tag == string(UpdatedAtTag) ||
					tag == string(CreatedAtTag) {
					continue
				}
				br.SelectParam[tag] = true
			}
			break
		}
		br.SelectParam[field] = true
	}
	return br
}

func (br *BaseRepository) Update(neo4jTransaction neo4j.Transaction, data interface{}) error {
	if br.BaseModel == nil {
		return errors.New("missing model")
	}
	if len(br.Condition) == 0 {
		return errors.New("missing condition")
	}
	input := make(map[string]interface{})
	//Check type
	typePointer := reflect.ValueOf(data).Kind()
	if typePointer != reflect.Ptr {
		return errors.New("invalid pointer type")
	}
	labelName := getStructName(data)
	acronymName := strings.ToLower(labelName[:1])
	matchQueryString := fmt.Sprintf(`MATCH (%s:%s) WHERE`, acronymName, labelName)
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
		input[field] = value
	}
	matchQueryString += fmt.Sprintf(`%s.deletedAt = ""`, acronymName)
	setString := ` SET`
	newValue := reflect.ValueOf(data).Elem()
	for i := 0; i < newValue.NumField(); i++ {
		// Get the field tag value
		tag := newValue.Type().Field(i).Tag.Get("neo4j")
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
		if len(br.SelectParam) != 0 {
			if br.SelectParam[tag] {
				timeValue, ok := newValue.Field(i).Interface().(time.Time)
				if ok {
					setString += fmt.Sprintf(` %s.%s= datetime($%s),`, acronymName, tag, tag)
					input[tag] = timeValue.Format(time.RFC3339)
				} else {
					setString += fmt.Sprintf(` %s.%s= $%s,`, acronymName, tag, tag)
					input[tag] = newValue.Field(i).Interface()
				}
			}
		} else {
			if !newValue.Field(i).IsZero() {
				timeValue, ok := newValue.Field(i).Interface().(time.Time)
				if ok {
					setString += fmt.Sprintf(` %s.%s= datetime($%s),`, acronymName, tag, tag)
					input[tag] = timeValue.Format(time.RFC3339)
				} else {
					setString += fmt.Sprintf(` %s.%s= $%s,`, acronymName, tag, tag)
					input[tag] = newValue.Field(i).Interface()
				}
			}
		}
	}
	setString += fmt.Sprintf(` %s.updatedAt= datetime($updatedAt)`, acronymName)
	input["updatedAt"] = time.Now().Format(time.RFC3339)
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
	queryString := matchQueryString + setString + returnQueryString
	fmt.Println(queryString)
	fmt.Println(input)
	result, err := neo4jTransaction.Run(queryString, input)
	if err != nil {
		fmt.Printf("error when update %s, error: %v", labelName, err)
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
	newValue.Set(value)
	return nil
}
