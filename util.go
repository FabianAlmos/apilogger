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

func constructLogFormat(level logLevel, msg string) string {
	time := time.Now().Format("2006-01-02-15:04:05 MST")
	if level == info || level == warn {
		return fmt.Sprintf("[ %s  - %s ]:\t", level, time) + msg
	}
	return fmt.Sprintf("[ %s - %s ]:\t", level, time) + msg
}
