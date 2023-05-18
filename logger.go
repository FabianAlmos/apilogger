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
)

// colors
const (
	_RED         escCode = 9
	_YELLOW      escCode = 11
	_MAGENTA     escCode = 13
	_AQUA        escCode = 14
	_DARKEST_RED escCode = 52
)

type RGBCode struct {
	R, G, B escCode
}

type logLevel string

const (
	info   logLevel = "INFO"
	warn   logLevel = "WARN"
	error_ logLevel = "ERROR"
	fatal  logLevel = "FATAL"
	debug  logLevel = "DEBUG"
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
	style := optionBuilder(code, dim, rgb.R, rgb.G, rgb.B)
	return style + msg + resetOption
}

func constructLogFormat(level logLevel, msg string) string {
	time := time.Now().Format("2006-01-02-15:04:05 MST")
	return fmt.Sprintf("[ %s\t- %s ]:\t", level, time) + msg
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

func Fatal(msg string) {
	message := constructLogFormat(fatal, msg)
	fmt.Println(msgBuilder(message, foreground, slowblink, _DARKEST_RED))
}

func Debug(msg string) {
	message := constructLogFormat(debug, msg)
	fmt.Println(msgBuilder(message, foreground, slowblink, _MAGENTA))
}

func DebugRGB(msg string, rgb RGBCode) {
	message := constructLogFormat(debug, msg)
	fmt.Println(msgBuilderRGB(message, foreground, rgb))
}
