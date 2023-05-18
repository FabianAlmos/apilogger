package logger

import (
	"fmt"
	"time"
)

type escCode uint8

// option codes
const (
	reset      escCode = 0
	dim        escCode = 2
	slowblink  escCode = 5
	foreground escCode = 38
	background escCode = 48
)

// colors
const (
	_RED    escCode = 9
	_YELLOW escCode = 11
	_AQUA   escCode = 14
)

type RGBCode struct {
	r, g, b escCode
}

type logLevel string

const (
	info   logLevel = "INFO"
	warn   logLevel = "WARN"
	error_ logLevel = "ERROR"
)

func optionBuilder(codes ...escCode) string {
	res := "\033["
	for index, code := range codes {
		if index < len(codes)-1 {
			res += fmt.Sprintf("%d", code) + ";"
		} else {
			res += fmt.Sprintf("%d", code)
		}
	}
	res += "m"
	return res
}

func msgBuilder(msg string, codes ...escCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(codes...)
	return style + msg + resetOption
}

func msgBuilderRGB(msg string, code escCode, rgb RGBCode) string {
	resetOption := optionBuilder(reset)
	style := optionBuilder(code, dim, rgb.r, rgb.g, rgb.b)
	return style + msg + resetOption
}

func constructLogFormat(level logLevel, msg string) string {
	time := time.Now().Format("2006-01-02::15:04:05 MST")
	return fmt.Sprintf("[ %s - %s ]:\t", level, time) + msg
}

func Info(msg string) {
	message := constructLogFormat(info, msg)
	fmt.Println(msgBuilder(message, foreground, slowblink, _AQUA))
}

func Warn(msg string) {
	message := constructLogFormat(warn, msg)
	fmt.Println(msgBuilder(message, foreground, slowblink, _YELLOW))
}

func Error(msg string) {
	message := constructLogFormat(error_, msg)
	fmt.Println(msgBuilder(message, foreground, slowblink, _RED))
}
