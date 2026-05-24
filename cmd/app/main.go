package main

import (
	"fmt"
	"log"

	"github.com/bodyanich/lab1-tooling/internal"
	"github.com/bodyanich/lab1-tooling/internal/config"
	"github.com/bodyanich/lab1-tooling/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	appLogger, err := logger.New(cfg.App.Environment)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer func() {
		if err := appLogger.Sync(); err != nil {
			log.Printf("failed to sync logger: %v", err)
		}
	}()

	appLogger.Info(
		"application started",
		zap.String("name", cfg.App.Name),
		zap.String("version", cfg.App.Version),
		zap.String("environment", cfg.App.Environment),
		zap.String("server_host", cfg.Server.Host),
		zap.Int("server_port", cfg.Server.Port),
	)

	sum := internal.Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", sum)

	result, err := internal.Divide(10, 2)
	if err != nil {
		log.Fatalf("failed to divide numbers: %v", err)
	}

	fmt.Printf("10 / 2 = %d\n", result)

	number := 4
	fmt.Printf("Is %d even? %v\n", number, internal.IsEven(number))
}
