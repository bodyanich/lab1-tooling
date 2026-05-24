package logger

import "go.uber.org/zap"

// New creates a new application logger.
func New(environment string) (*zap.Logger, error) {
	if environment == "production" {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}
