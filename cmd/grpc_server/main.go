package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/repository/model"
	"github.com/vakhrushevk/chat-server-service/internal/repository/postgres"

	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/internal/config/env"

	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	chatRepository repository.ChatRepository
	chat_v1.UnimplementedChatV1Server
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

	srver := &server{chatRepository: rep}

	chat_v1.RegisterChatV1Server(srv, srver)
	log.Printf("server listening at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat - this method is intended to create a new chat.
// It takes context and a request for creating a chat *chat_v1.CreateChatRequest.
// returns a response in the form of the chat_v1.CreateChatResponse structure, containing the identifier of the created chat and an error
func (s *server) CreateChat(ctx context.Context, request *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	// temp
	chat := model.Chat{Name: request.ChatName}
	id, err := s.chatRepository.CreateChat(ctx, chat, request.IdUsers)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Create Chat: %v", id)

	return &chat_v1.CreateChatResponse{IdChat: id}, nil
}

// SendMessage - this method is intended to create a new mewssage,
// It takes context.Context. and a request for create message *chat_v1.SendMessageRequest.
func (s *server) SendMessage(ctx context.Context, request *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("Send Message: %v", request)
	message := model.Message{Text: request.Text, Sender: request.FromIdUser, IDChat: request.IdChat}
	err := s.chatRepository.SendMessage(ctx, message)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteChat - this method is intended to delete chat.
// It takes context and a request for deleting chat *chat_v1.DeleteChatRequest
func (s server) DeleteChat(ctx context.Context, request *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	err := s.chatRepository.DeleteChat(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
