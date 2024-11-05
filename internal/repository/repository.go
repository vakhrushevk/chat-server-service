package repository

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/repository/model"
)

// ChatRepository -
type ChatRepository interface {
	// CreateChat - Создает чат и заполняет его юзерами? возвращает id чата и ошибку
	CreateChat(ctx context.Context, chat model.RepoChat, userID []int64) (int64, error)
	// SendMessage - Создает сообщение, возвращает ошибку
	SendMessage(ctx context.Context, message model.RepoMessage) error
	// DeleteChat - Удалет чат из базы данных, возвращает ошибку
	DeleteChat(ctx context.Context, idChat int64) error
}
