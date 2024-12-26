package chatservice

import (
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/local-platform/pkg/db"
)

type serv struct {
	repositoy repository.ChatRepository
	txManager db.TxManager
}

// New - TODO: Add description
func New(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{repositoy: chatRepository, txManager: txManager}
}
