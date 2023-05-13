package logger

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

// log out codes
const (
	LOCAL    logOutCode = 0x01
	TERMINAL logOutCode = 0x02
	NETLOG   logOutCode = 0x04
	ALL      logOutCode = 0x08
)
