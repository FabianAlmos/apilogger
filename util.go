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

func addData(data ...any) string {
	res := ""
	for index, d := range data {
		if index < len(data)-1 {
			res += "\t" + fmt.Sprintln(d)
		} else {
			res += "\t" + fmt.Sprint(d)
		}
	}
	if res != "" {
		return "\n" + res
	}
	return ""
}

func constructLogFormat(level logLevel, msg string, data ...any) string {
	time := time.Now().Format("2006-01-02-15:04:05 MST")
	if level == info || level == warn {
		return fmt.Sprintf("[ %s  - %s ]:\t", level, time) + msg + addData(data...)
	}
	return fmt.Sprintf("[ %s - %s ]:\t", level, time) + msg + addData(data...)
}
