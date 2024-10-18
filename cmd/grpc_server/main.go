package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/internal/config/env"
	"log"
	"net"

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
	chat_v1.RegisterChatV1Server(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat - this method is intended to create a new chat.
// It takes context and a request for creating a chat *chat_v1.CreateChatRequest.
// returns a response in the form of the chat_v1.CreateChatResponse structure, containing the identifier of the created chat and an error
func (s server) CreateChat(_ context.Context, request *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	fmt.Printf("Create Chat: %v", request)
	return &chat_v1.CreateChatResponse{IdChat: 1}, nil
}

// SendMessage - this method is intended to create a new mewssage,
// It takes context.Context. and a request for create message *chat_v1.SendMessageRequest.
func (s server) SendMessage(_ context.Context, request *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("Send Message: %v", request)
	return nil, nil
}

// DeleteChat - this method is intended to delete chat.
// It takes context and a request for deleting chat *chat_v1.DeleteChatRequest
func (s server) DeleteChat(_ context.Context, request *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete Chat: %v", request)
	return nil, nil
}
