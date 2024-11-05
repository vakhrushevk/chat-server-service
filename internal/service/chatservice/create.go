package chatservice

import (
	"context"
	"log"

	"github.com/vakhrushevk/chat-server-service/internal/service/model"
	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
)

// CreateChat - TODO: Add description
func (s *serv) CreateChat(ctx context.Context, chat *model.ServiceChat) (int64, error) {
	// TODO: Обработка ошибок из репозитория
	log.Println("Я тут был")
	return s.repositoy.CreateChat(ctx, *converter.FromChatToRepo(chat), chat.UserID)
}
