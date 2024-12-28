package chatservice

import (
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/local-platform/pkg/db"
)

type serv struct {
	repository repository.ChatRepository
	txManager  db.TxManager
}

// New - creates a new chat level service
func New(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{repository: chatRepository, txManager: txManager}
}
