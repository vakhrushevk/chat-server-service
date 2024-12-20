package chat

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat - this method is intended to delete chatservice.
// It takes context and a request for deleting chatservice *chat_v1.DeleteChatRequest
func (iml *Implementation) DeleteChat(ctx context.Context, request *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	err := iml.chatService.DeleteChat(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
