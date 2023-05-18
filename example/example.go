package main

import (
	logger "github.com/FabianAlmos/apilogger"
)

func init() {
	// Configuration for the logger must be written in the init function!
	// Configuration can be set by assigning a new 'logger.LoggerConfig' struct to the global variable 'logger.Config'.
	// Or by setting the fields of 'logger.Config' directly like 'logger.Config.Out = logger.ALL'.
	// 'Out' in 'logger.LoggerConfig' is the target output for logging.
	// Outputs can be: 'FILE', 'TERMINAL', 'ALL'

	// 'FILE' option creates a 'logs' directory in the scope of the entry point where the logger is started,
	// and the logging functions will write to a log file in it, every run a new log file is created.

	// 'TERMINAL' option logs to the terminal directly.

	// 'ALL' option combines the behaviour of all the other options (called 'ALL' incase of adding new logging targets).

	// The field 'Server' is for future use currently it doesn't do anything right now!

	logger.Config = &logger.LoggerConfig{
		Out:    logger.ALL,
		Server: "",
	}

	// After setting the configurations 'logger.Start()' must be called to start logging,
	// otherwise the logging with option 'FILE' won't work!
	logger.Start()
}

func main() {
	// Adviced to call 'logger.Stop()' with 'defer' here on the first line of 'main'.
	// If the logger isn't stopped than the 'FILE' option won't work!
	defer logger.Stop()

	logger.Info("INFO")
	logger.Warn("WARN")
	logger.Error("ERROR")
	logger.Fatal("FATAL")
	logger.Debug("DEBUG")
	logger.DebugRGB("DEBUG_RGB", logger.RGBCode{R: 3, G: 252, B: 107})
}
