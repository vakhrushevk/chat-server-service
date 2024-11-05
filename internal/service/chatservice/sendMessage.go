package chatservice

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/service/model"
	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
)

// SendMessage - Send Message! OK? ok
func (s *serv) SendMessage(ctx context.Context, message *model.ServiceMessage) error {
	// TODO: CheckError

	return s.repositoy.SendMessage(ctx, *converter.FromServiceMessageToRepoMessage(message))
}
