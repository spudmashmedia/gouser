package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spudmashmedia/gouser/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("Config loading failed", "error", err))
	}

	opt := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	slog.SetDefault(logger)

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error(
			"Server failed to start",
			"error", err)
		os.Exit(1)
	}
}
