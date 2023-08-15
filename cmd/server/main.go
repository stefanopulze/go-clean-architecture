package main

import (
	"go-clean-architecture/internal/api"
	"go-clean-architecture/internal/config"
	"go-clean-architecture/internal/db"
	"go-clean-architecture/internal/logging"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Load logger
	logger := logging.Init(cfg.Log)
	logger.Info("Starting API Server")

	// Connect database
	if err := db.MustConnect(cfg.Database); err != nil {
		logger.Error("Cannot connect to database", "error", err.Error())
		panic(err)
	}
	defer db.Close()

	api.Start(cfg)
}
