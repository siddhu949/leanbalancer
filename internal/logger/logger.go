package logger

import (
	"go.uber.org/zap"
)

// Logger instance
var Logger *zap.Logger

// Initialize Logger
func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // Use zap.NewDevelopment() for debug mode
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer Logger.Sync()
}

// Get Logger instance
func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
