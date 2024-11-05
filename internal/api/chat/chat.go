package chat

import (
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
)

// Implementation - TODO: Add description
type Implementation struct {
	chat_v1.UnimplementedChatV1Server

	chatService service.ChatService
}

// NewImplementation - TODO: add description
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{chatService: chatService}
}
