package main

import (
	logger "github.com/FabianAlmos/apilogger"
)

func init() {
	// Configuration for the logger must be written in the init function, before calling 'logger.Start()'!
	// Configuration can be set by assigning a new 'logger.LoggerConfig' struct to the global variable 'logger.Config'.
	// Or by setting the fields of 'logger.Config' directly like 'logger.Config.Out = logger.ALL'.

	// 'InDebug' option controls whether the logger can use the 'logger.Debug' and 'logger.DebugRGB' logging function.
	// If 'InDebug' is set to false than debugging functions will give a warn level message saying 'NOT IN DEBUG MODE!'

	// 'Out' in 'logger.LoggerConfig' is the target output for logging.
	// Outputs can be: 'FILE', 'TERMINAL', 'ALL'

	// 'logger.FILE' option creates a 'logs' directory in the scope of the entry point where the logger is started,
	// and the logging functions will write to a log file in it, every run a new log file is created.

	// 'logger.TERMINAL' option logs to the terminal directly.

	// 'logger.ALL' option combines the behaviour of all the other options (called 'ALL' incase of adding new logging targets).

	// The field 'Server' is for future use, currently it doesn't do anything!

	// If the Config for the logger isn't set than the logger will use the base configurations, which are the following:
	//
	// InDebug: false
	// Out: logger.TERMINAL
	// Server: ""

	logger.Config = &logger.LoggerConfig{
		InDebug: true,
		Out:     logger.ALL,
		Server:  "",
	}

	// After setting the configurations 'logger.Start()' must be called to start logging,
	// otherwise the logging with option 'logger.FILE' won't work!
	logger.Start()
}

func main() {
	// Adviced to call 'logger.Stop()' with 'defer' here on the first line of 'main'.
	// If the logger isn't stopped than the 'FILE' option won't work!
	defer logger.Stop()

	logger.Info("INFO", "dataInfo")
	logger.Warn("WARN", "dataWarn", "dataWarn1")
	logger.Error("ERROR", "dataErr")
	logger.Debug("DEBUG")
	logger.DebugRGB("DEBUG_RGB", logger.RGBCode{R: 3, G: 252, B: 107})

	// 'logger.Fatal' is the only logging funtion that will stop the execution of a goroutine
	logger.Fatal("FATAL", "panic value", "data", "data1", "data2")
}
