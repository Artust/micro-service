package base_repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (br *BaseRepository) CreateMany(neo4jTransaction neo4j.Transaction, data interface{}) error {
	input := make(map[string]interface{})
	//Check type
	typePointer := reflect.ValueOf(data).Kind()
	if typePointer != reflect.Ptr {
		return errors.New("invalid pointer type")
	}
	typePointerSlice := reflect.TypeOf(data).Elem().Kind()
	if typePointerSlice != reflect.Slice {
		return errors.New("invalid slice")
	}
	labelName := getStructName(data)
	acronymName := strings.ToLower(labelName[:1])
	createQueryString := `CREATE `
	returnQueryString := ` RETURN `
	sliceValue := reflect.ValueOf(data).Elem()
	if sliceValue.Len() == 0 {
		return errors.New("invalid slice")
	}
	for i := 0; i < sliceValue.Len(); i++ {
		createElementString := fmt.Sprintf(`(%s%d: %s{`, acronymName, i, labelName)
		returnElementString := fmt.Sprintf(`id(%s%d), `, acronymName, i)
		value := sliceValue.Index(i)
		for j := 0; j < value.NumField(); j++ {
			// Get the field tag value
			tag := value.Type().Field(j).Tag.Get("neo4j")
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
			createElementString += fmt.Sprintf(`%s:$%s%d,`, tag, tag, i)
			returnElementString += fmt.Sprintf(`%s%d.%s, `, acronymName, i, tag)
			input[fmt.Sprintf(`%s%d`, tag, i)] = value.Field(j).Interface()
		}
		createElementString += fmt.Sprintf(`createdAt: datetime($createdAt%d), `, i)
		createElementString += fmt.Sprintf(`updatedAt:datetime($updatedAt%d), `, i)
		createElementString += `deletedAt:""})`
		returnElementString += fmt.Sprintf(`%s%d.createdAt, `, acronymName, i)
		returnElementString += fmt.Sprintf(`%s%d.updatedAt `, acronymName, i)
		input[fmt.Sprintf(`createdAt%d`, i)] = time.Now().Format(time.RFC3339)
		input[fmt.Sprintf(`updatedAt%d`, i)] = time.Now().Format(time.RFC3339)
		if i != sliceValue.Len()-1 {
			// handle comma
			createElementString += `,`
			returnElementString += `,`
		}
		createQueryString += createElementString
		returnQueryString += returnElementString
	}
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
		fmt.Printf("error when create %s, error: %v", labelName, err)
		return err
	}
	destinationValue := reflect.ValueOf(data)
	output := reflect.MakeSlice(destinationValue.Elem().Type(), sliceValue.Len(), sliceValue.Len())
	value := sliceValue.Index(0)
	for j := 0; j < sliceValue.Len(); j++ {
		for i := 0; i < value.NumField(); i++ {
			// Get the field tag value
			tag := value.Type().Field(i).Tag.Get("neo4j")
			if tag == string(IdTag) {
				createdValue, ok := record.Get(fmt.Sprintf("id(%s%d)", acronymName, j))
				if ok {
					value.Field(i).Set(reflect.ValueOf(createdValue))
				}
				continue
			}
			if tag == string(DeletedAtTag) {
				continue
			}
			createdValue, ok := record.Get(fmt.Sprintf("%s%d.%s", acronymName, j, tag))
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
		output.Index(j).Set(value)
	}
	destinationValue.Elem().Set(output)
	return nil
}
