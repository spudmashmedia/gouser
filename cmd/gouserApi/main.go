package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spudmashmedia/gouser/internal/api"
	"github.com/spudmashmedia/gouser/internal/config"
	"github.com/spudmashmedia/gouser/internal/logger"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("Config loading failed %s", err))
	}

	logger.RegisterLogger(cfg.Logger.LogLevel)

	app := api.NewApplication(cfg)

	if err := app.Run(app.Mount()); err != nil {
		slog.Error(
			"Server failed to start",
			"error", err)
		os.Exit(1)
	}
}
