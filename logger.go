package logger

import (
	"fmt"
	"os"
	"time"
)

var Config LoggerConfig = LoggerConfig{
	Out:    TERMINAL,
	Server: "",
}

var logFile *os.File

func init() {
	fileName := "logs/log-" + time.Now().Format("2006-01-02-15:04:05 MST") + ".txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	logFile = file
}

func Info(msg string) string {
	message := constructLogFormat(info, msg)
	return msgBuilder(message, foreground, slowblink, _AQUA)
}

func Warn(msg string) string {
	message := constructLogFormat(warn, msg)
	return msgBuilder(message, foreground, slowblink, _YELLOW)
}

func Error(msg string) string {
	message := constructLogFormat(error_, msg)
	return msgBuilder(message, foreground, slowblink, _RED)
}

func Fatal(msg string) string {
	message := constructLogFormat(fatal, msg)
	return msgBuilder(message, foreground, slowblink, _DARKEST_RED)
}

func Debug(msg string) string {
	message := constructLogFormat(debug, msg)
	return msgBuilder(message, foreground, slowblink, _MAGENTA)
}

func DebugRGB(msg string, rgb RGBCode) string {
	message := constructLogFormat(debug, msg)
	return msgBuilderRGB(message, foreground, rgb)
}

func StopLogging() {
	logFile.Close()
}
