package logger

import (
	"fmt"
)

var Config LoggerConfig = LoggerConfig{
	Out:    TERMINAL,
	Server: "",
}

func Info(msg string) string {
	message := constructLogFormat(info, msg)
	builtMsg := msgBuilder(message, foreground, slowblink, _AQUA)
	fmt.Println(builtMsg)
	return builtMsg
}

func Warn(msg string) string {
	message := constructLogFormat(warn, msg)
	builtMsg := msgBuilder(message, foreground, slowblink, _YELLOW)
	fmt.Println(builtMsg)
	return builtMsg
}

func Error(msg string) string {
	message := constructLogFormat(error_, msg)
	builtMsg := msgBuilder(message, foreground, slowblink, _RED)
	fmt.Println(builtMsg)
	return builtMsg
}

func Fatal(msg string) string {
	message := constructLogFormat(fatal, msg)
	builtMsg := msgBuilder(message, foreground, slowblink, _DARKEST_RED)
	fmt.Println(builtMsg)
	return builtMsg
}

func Debug(msg string) string {
	message := constructLogFormat(debug, msg)
	builtMsg := msgBuilder(message, foreground, slowblink, _MAGENTA)
	fmt.Println(builtMsg)
	return builtMsg
}

func DebugRGB(msg string, rgb RGBCode) string {
	message := constructLogFormat(debug, msg)
	builtMsg := msgBuilderRGB(message, foreground, rgb)
	fmt.Println(builtMsg)
	return builtMsg
}
