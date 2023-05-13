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
	_LOCAL    logOutCode = 0x01
	_TERMINAL logOutCode = 0x02
	_NETLOG   logOutCode = 0x04
	_LOC_TERM logOutCode = 0x08
	_LOC_NET  logOutCode = 0x10
	_TERM_NET logOutCode = 0x20
	_ALL      logOutCode = 0x40
)
