package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func LoadLoggerConfig(path string) (*zap.SugaredLogger, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("logger config file does not exist: %s", path)
	}
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg zap.Config
	if err = json.Unmarshal(configData, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	logger, err := cfg.Build()

	if err != nil {
		return nil, fmt.Errorf("failed to build logger from config %q: %w", path, err)
	}
	zap.ReplaceGlobals(logger)

	return logger.Sugar(), nil
}
