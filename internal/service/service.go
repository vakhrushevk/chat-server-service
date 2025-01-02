package service

import (
	"context"
	"github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
)

type ChatService interface {
	CreateChat(ctx context.Context, chat *serviceLevelModel.ChatInfo) (int64, error)
	DeleteChat(ctx context.Context, idChat int64) error
	AddChatMember(ctx context.Context, chat *serviceLevelModel.ChatMemberInfo) error
	RemoveChatMember(ctx context.Context, chat *serviceLevelModel.ChatMemberInfo) error
	ListChatsByIdUser(ctx context.Context, UserID int64) ([]*serviceLevelModel.Chat, error)
}
