package logger

type logOutput uint8

type LoggerConfig struct {
	InDebug bool
	Out     logOutput
	Server  string
}
