package main

import (
	"context"
	"flag"
	"github.com/vakhrushevk/local-platform/pkg/logger"
	"log"

	"github.com/vakhrushevk/chat-server-service/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	logger.Init()
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
