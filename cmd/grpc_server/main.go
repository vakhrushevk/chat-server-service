package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/vakhrushevk/chat-server-service/internal/api/chat"
	"github.com/vakhrushevk/chat-server-service/internal/service/chatservice"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vakhrushevk/chat-server-service/internal/repository/postgres"

	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/internal/config/env"

	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpcConfig: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load pgConfig: %v", err)
	}
	_ = pgConfig

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}

	rep := postgres.NewRepository(pool)
	chatServ := chatservice.New(rep)
	serv := chat.NewImplementation(chatServ)
	chat_v1.RegisterChatV1Server(srv, serv)
	log.Printf("server listening at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
