package service

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/service/model"
)

// ChatService - TODO: add description
type ChatService interface {
	// CreateChat - Создаем чат,
	CreateChat(ctx context.Context, chat *model.ServiceChat) (int64, error)
	// SendMessage - Создает сообщение, возвращает ошибку
	SendMessage(ctx context.Context, message *model.ServiceMessage) error
	// DeleteChat - Удалем чат
	DeleteChat(ctx context.Context, idChat int64) error
}
