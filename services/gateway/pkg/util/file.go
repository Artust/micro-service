package util

import (
	"fmt"
	"strings"
	"time"
)

func GenerateFileName(fileNameInput string, name string) (fileNameOutput string, format string) {
	t := time.Now()
	formatTime := fmt.Sprintf("%d%02d%02dT%02d%02d%02d.%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	dotName := strings.LastIndex(fileNameInput, ".")
	fileNameOutput = fmt.Sprintf("%v-%v%v", name, formatTime, fileNameInput[dotName:])
	format = strings.ToUpper(fileNameInput[dotName:])
	return fileNameOutput, format
}
