package chatservice

import (
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/service"
)

type serv struct {
	repositoy repository.ChatRepository
}

// New - TODO: Add description
func New(chatRepository repository.ChatRepository) service.ChatService {
	return &serv{repositoy: chatRepository}
}
