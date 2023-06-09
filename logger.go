package logger

import (
	"fmt"
	"io/fs"
	"os"
	"time"
)

func Start() {
	if Config.Out&FILE == FILE || Config.Out&ALL == ALL {
		var perms fs.FileMode = 0755
		_ = os.Mkdir("logs", perms)
		fileName := "logs/log-" + time.Now().Format("2006-01-02-15-04-05MST") + ".txt"
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, perms)
		if err != nil {
			fmt.Println(err)
		}
		logFile = file
	}
}

func Stop() {
	logFile.Close()
}

func Info(msg string, data ...any) string {
	message := constructLogFormat(info, msg, data...)
	return msgBuilder(message, foreground, slowblink, _AQUA)
}

func Warn(msg string, data ...any) string {
	message := constructLogFormat(warn, msg, data...)
	return msgBuilder(message, foreground, slowblink, _YELLOW)
}

func Error(msg string, data ...any) string {
	message := constructLogFormat(error_, msg, data...)
	return msgBuilder(message, foreground, slowblink, _RED)
}

func Fatal(msg string, p any, data ...any) string {
	message := constructLogFormat(fatal, msg, data...)
	msgBuilder(message, foreground, slowblink, _DARKEST_RED)
	panic(p)
}

func Debug(msg string, data ...any) string {
	if Config.InDebug {
		message := constructLogFormat(debug, msg, data...)
		return msgBuilder(message, foreground, slowblink, _MAGENTA)
	} else {
		return Warn("NOT IN DEBUG MODE!")
	}
}

func DebugRGB(msg string, rgb RGBCode, data ...any) string {
	if Config.InDebug {
		message := constructLogFormat(debug, msg, data...)
		return msgBuilderRGB(message, foreground, rgb)
	} else {
		return Warn("NOT IN DEBUG MODE!")
	}
}
