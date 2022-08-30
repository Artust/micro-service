package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) Model(model interface{}) *BaseRepository {
	newBr := CreateBaseRepository(br.Neo4j)
	newBr.BaseModel = model
	return newBr
}

func (br *BaseRepository) Where(key string, value interface{}) *BaseRepository {
	if br.Condition == nil {
		br.Condition = make(map[string]interface{})
	}
	data := reflect.ValueOf(value).Interface()
	timeValue, ok := data.(time.Time)
	if ok {
		br.Condition[key] = timeValue
		return br
	}
	br.Condition[key] = value
	return br
}

func (br *BaseRepository) Skip(skip int64) *BaseRepository {
	br.SkipParam = 0
	if skip == 0 {
		return br
	}
	br.SkipParam = skip
	return br
}

func (br *BaseRepository) Limit(limit int64) *BaseRepository {
	if limit == 0 {
		return br
	}
	br.LimitParam = limit
	return br
}

func (br *BaseRepository) Order(order string) *BaseRepository {
	splitString := strings.Split(order, " ")
	if len(splitString) < 2 {
		return br
	}
	orderType := "DESC"
	if strings.EqualFold(splitString[1], "ASC") {
		orderType = "ASC"
	}
	br.OrderParam = fmt.Sprintf(`%s %s`, splitString[0], orderType)
	return br
}

func (br *BaseRepository) Find(neo4jTransaction neo4j.Transaction, destination interface{}) error {
	if br.BaseModel == nil {
		return errors.New("missing model")
	}
	//Check type
	typePointer := reflect.ValueOf(destination).Kind()
	if typePointer != reflect.Ptr {
		return errors.New("invalid pointer type")
	}
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
			timeValue, ok := value.(time.Time)
			if ok {
				matchQueryString += fmt.Sprintf(` %s.%s %s datetime($%s) AND `, acronymName, field, operator, field)
				conditionQuery[field] = timeValue.Format(time.RFC3339)
				continue
			}
			matchQueryString += fmt.Sprintf(` %s.%s %s $%s AND `, acronymName, field, operator, field)
		}
		conditionQuery[field] = value
	}
	matchQueryString += fmt.Sprintf(` %s.deletedAt = ""`, acronymName)
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
	records, err := result.Collect()
	if err != nil {
		fmt.Printf("error when get %s, error: %v", labelName, err)
		return err
	}
	destinationValue := reflect.ValueOf(destination)
	output := reflect.MakeSlice(destinationValue.Elem().Type(), len(records), len(records))
	for j, record := range records {
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
		valuePointer := reflect.New(value.Type())
		valuePointer.Elem().Set(value)
		output.Index(j).Set(valuePointer)
	}
	destinationValue.Elem().Set(output)
	return nil
}
