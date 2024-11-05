package chat

import (
	"context"
	"fmt"

	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
)

// CreateChat - this method is intended to create a new chatservice.
// It takes context and a request for creating a chatservice *chat_v1.CreateChatRequest.
// returns a response in the form of the chat_v1.CreateChatResponse structure, containing the identifier of the created chatservice and an error
func (iml *Implementation) CreateChat(ctx context.Context, request *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {

	id, err := iml.chatService.CreateChat(ctx, converter.FromDescToChat(request))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Create Chat: %v", id)

	return &chat_v1.CreateChatResponse{IdChat: id}, nil
}
