package util

import (
	"fmt"
	"time"
)

func FormatRFC3339(input interface{}) string {
	if len(fmt.Sprintf("%v", input)) > 0 {
		time := input.(time.Time).Format(time.RFC3339)
		return fmt.Sprintf("%v", time)
	}
	return ""
}

func ParseTime(input string) time.Time {
	timeValue, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Time{}
	}
	return timeValue
}
