package logger

// logger outputs
const (
	TERMINAL logOutput = 0x01
	FILE     logOutput = 0x02
	//NETLOG   logOutput = 0x04 //TODO: Add capability to log into a server for an api for an example
	ALL logOutput = 0x08
)

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

// levels
const (
	info   logLevel = "INFO"
	warn   logLevel = "WARN"
	error_ logLevel = "ERROR"
	fatal  logLevel = "FATAL"
	debug  logLevel = "DEBUG"
)
