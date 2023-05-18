package logger

import (
	"fmt"
	"time"
)

type escCode uint8
type logLevel string
type RGBCode struct {
	R, G, B escCode
}

func addData(data string) string {
	if data != "" {
		return "\n\t" + data
	}
	return ""
}

func constructLogFormat(level logLevel, msg string, data string) string {
	time := time.Now().Format("2006-01-02-15:04:05 MST")
	if level == info || level == warn {
		return fmt.Sprintf("[ %s  - %s ]:\t", level, time) + msg + addData(data)
	}
	return fmt.Sprintf("[ %s - %s ]:\t", level, time) + msg + addData(data)
}
