package logger

type logOutput uint8

type LoggerConfig struct {
	Out    logOutput
	server string
}
