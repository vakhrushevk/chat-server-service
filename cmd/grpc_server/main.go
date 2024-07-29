package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = ":50051"

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
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

// CreateChat - Создать чат
//
//nolint:revive
func (s *server) CreateChat(ctx context.Context, request *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	fmt.Println("Received Create request")
	if request.GetUsernames() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Usernames not provided")
	}
	for id, user := range request.GetUsernames() {
		fmt.Printf("%d) %s - add to chat", id+1, user)
	}
	return &chat_v1.CreateChatResponse{Id: 1}, nil
}

// SendMessage - Отправить сообщение
//
//nolint:revive
func (s *server) SendMessage(ctx context.Context, request *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Println("Received SendMessage request")
	fmt.Printf("\nFrom: %s\n Text: %s\n Time: %s\n", request.GetFrom(), request.GetText(), request.GetTimestamp().AsTime())

	return &emptypb.Empty{}, nil
}

// DeleteChat - Удалить чат
//
//nolint:revive
func (s *server) DeleteChat(ctx context.Context, request *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	fmt.Println("Received Delete request:")
	fmt.Println(request.GetId())

	return &emptypb.Empty{}, nil
}
