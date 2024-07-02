package main

import (
	"github.com/rafaamandio/gopportunites/config"
	"github.com/rafaamandio/gopportunites/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization erro: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
