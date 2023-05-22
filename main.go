package main

import (
	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
)

func main() {
	config := config.NewDefaultConfig()
	if err := config.LoadFromToml("config/testdata/sample.toml"); err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(config.DevMode)
	if err != nil {
		panic(err)
	}

	logger.Info("Go server boilerplate starting...")
}
