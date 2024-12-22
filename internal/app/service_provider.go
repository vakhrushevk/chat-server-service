package app

import (
	"context"
	db "github.com/vakhrushevk/chat-server-service/internal/client"
	"github.com/vakhrushevk/chat-server-service/internal/client/db/pg"
	"log"

	"github.com/vakhrushevk/chat-server-service/internal/api/chat"
	"github.com/vakhrushevk/chat-server-service/internal/closer"
	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/internal/config/env"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/repository/postgres"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/chat-server-service/internal/service/chatservice"
)

type serviceProvider struct {
	pgConfig   config.PgConfig
	grpcConfig config.GRPCConfig

	dbClient db.Client

	chatRepository repository.ChatRepository
	chatService    service.ChatService

	chatImplementation *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DBClient() db.Client {
	if s.dbClient == nil {
		client, err := pg.New(context.Background(), s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}
		err = client.DB().Ping(context.Background())
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(client.Close)
		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("Failed to get pg config: %v", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		repo := postgres.NewChatRepository(s.DBClient())
		s.chatRepository = repo
	}
	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		serv := chatservice.New(s.ChatRepository(ctx))
		s.chatService = serv
	}
	return s.chatService
}

func (s *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if s.chatImplementation == nil {
		s.chatImplementation = chat.NewImplementation(s.ChatService(ctx))
	}
	return s.chatImplementation
}
