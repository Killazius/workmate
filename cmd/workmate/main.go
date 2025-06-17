package main

import (
	"github.com/Killazius/workmate/internal/config"
	"github.com/Killazius/workmate/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	log, err := logger.LoadLoggerConfig(cfg.LoggerPath)
	if err != nil {
		panic(err)
	}

}
