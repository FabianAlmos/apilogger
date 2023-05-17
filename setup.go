package logger

import "os"

var Config *LoggerConfig = new(LoggerConfig)
var logFile *os.File = new(os.File)

func init() {
	Config = &LoggerConfig{
		Out:    TERMINAL,
		Server: "",
	}
}
