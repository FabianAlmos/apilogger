package logger

import (
	"fmt"
)

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
