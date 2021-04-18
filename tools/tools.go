package tools

import (
	"fmt"
	"time"
)

//BytesToString converter
func BytesToString(data []byte) string {
	return string(data[:])
}

//GetStrDateTime Function
func GetStrDateTime() string {
	dt := time.Now().Format("2006-02-01 15:04:05")
	msg := fmt.Sprintf("%-19s", dt)
	return msg
}

//GetStrtimestamp Function
func GetStrtimestamp() string {
	dt := time.Now().Format("20060201150405.9999")
	msg := fmt.Sprintf("%s", dt)
	return msg
}
