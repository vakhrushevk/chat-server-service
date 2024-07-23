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

//nolint:revive
func (s *server) Create(ctx context.Context, request *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	fmt.Println("Received Create request")
	if request.GetUsernames() != nil {
		for id, user := range request.GetUsernames() {
			fmt.Println(id+1, user, " - aadd to chat")
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Usernames not provided")
	}

	return &chat_v1.CreateResponse{Id: 1}, nil
}

//nolint:revive
func (s *server) SendMessage(ctx context.Context, request *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Println("Received SendMessage request")
	fmt.Printf("\nFrom: %s\n Text: %s\n Time: %s\n", request.GetFrom(), request.GetText(), request.GetTimestamp().AsTime())

	return &emptypb.Empty{}, nil
}

//nolint:revive
func (s *server) Delete(ctx context.Context, request *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("Received Delete request:")
	fmt.Println(request.GetId())

	return &emptypb.Empty{}, nil
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
