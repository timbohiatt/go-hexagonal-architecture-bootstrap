package main

import (
	"log"
	"persona/cmd/config"
	"persona/internal/adapters/app/api"
	"persona/internal/adapters/core"
	"persona/internal/adapters/primary/http"
	"persona/internal/ports"
	"sync"
)

// Main

func main() {

	// Setup Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("Configuration Loading Failed - Exiting")
		log.Fatal(err)

	}

	// Validate Configuration
	err = cfg.Validate()
	if err != nil {
		log.Println("Configuration Validation Failed - Exiting")
		log.Fatal(err)
	}

	// Ports
	// Primary Ports
	var httpAdapter ports.HTTPPort
	// App Ports
	var apiAdapter ports.APIPort
	// Core Ports
	var coreAdapter ports.CorePort

	// Adapter Configuration
	// Core Adapters
	coreAdapter = core.NewAdapter()
	// App Adapters
	apiAdapter = api.NewAdapter(coreAdapter)
	// Primary Adapters
	httpAdapter = http.NewAdapter(apiAdapter)

	// Kick off Primary 
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run HTTP Server
	if cfg.HTTPServer.Enabled {
		log.Println("HTTP Server Enabled")
		wg.Add(1)
		go httpAdapter.Run(
			cfg.HTTPServer.Host,
			cfg.HTTPServer.Port,
			&wg)
	} else {
		log.Println("HTTP Server Disabled")
	}
	// Wait for all services to Finish Running
	wg.Wait()

}
