package main

import (
	"context"
	"flag"
	"log"

	"github.com/vakhrushevk/chat-server-service/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {

	flag.Parse()
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
