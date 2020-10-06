package app

import (
	"fmt"
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Configure Logging
func setupLogger() {
	c := config.C.Log
	logger.SetLevel(c.Level)
	logger.SetFormatter(c.Format)

	// TODO: default logs write to stderr
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		// TODO: write logs to logs file
		logger.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
		fmt.Printf("save logs to: %v\n", LOG_FILE_LOCATION)
	}
}
