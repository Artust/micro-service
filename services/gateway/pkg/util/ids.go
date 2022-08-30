package util

import (
	"strconv"
	"strings"
)

func GenerateIds(input string) []int64 {
	var ids []int64
	if input != "" {
		input := strings.Split(input, ",")
		for _, id := range input {
			id = strings.Trim(id, " ")
			id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return nil
			}
			ids = append(ids, id)
		}
	}
	return ids
}
