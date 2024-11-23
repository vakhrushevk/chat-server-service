package app

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
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

	pgPool *pgxpool.Pool

	chatRepository repository.ChatRepository
	chatService    service.ChatService

	chatImplementation *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())

		fmt.Println("я тут")
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
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
		repo := postgres.NewChatRepository(s.PgPool(ctx))
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
