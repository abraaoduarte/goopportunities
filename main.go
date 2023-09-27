package main

import (
	"github.com/abraaoduarte/goopportunities/config"
	"github.com/abraaoduarte/goopportunities/router"
)

var (
	logger *config.Logger
)
func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()

}

