package logging

import (
	"arp147/flags"
	"arp147/storage"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

// Stderr is a standard logger prints out to stderr
var Stderr = log.New(os.Stderr, "", log.LstdFlags)

// Stdout is a standard logger that prints to stdout
var Stdout = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	// If we're in debug mode we want to log everything to the console, but if
	// we're not in debug mode, we want to log everything to files.
	if !*flags.Debug {
		// We'll store everything in a "logs" folder in the storage directory.
		logs := storage.Collection.NewStorage("logs")

		// Informative information will be logged to "stdout.log"
		Stdout.SetOutput(&lumberjack.Logger{
			Filename: logs.GetPath() + "/stdout.log",
			MaxSize:  10,
		})

		// Errors will be logged to "stderr.log"
		Stderr.SetOutput(&lumberjack.Logger{
			Filename: logs.GetPath() + "/stderr.log",
			MaxSize:  10,
		})
	}
}
