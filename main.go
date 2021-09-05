package main

import (
	"cms-api/config"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	config.UpdateConfigsFor(env)
	// server.StartNewServer()
}
