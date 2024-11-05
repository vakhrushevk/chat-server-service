package chat

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage - this method is intended to create a new mewssage,
// It takes context.Context. and a request for create message *chat_v1.SendMessageRequest.
func (iml *Implementation) SendMessage(ctx context.Context, request *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	err := iml.chatService.SendMessage(ctx, converter.FromDescToServiceMessage(request))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
