package main

import (
	"cms-api/config"
	"cms-api/server"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	config.UpdateConfigsFor(env)
	server.StartNewServer(env)
}
