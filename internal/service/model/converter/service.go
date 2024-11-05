package converter

import (
	modelRepo "github.com/vakhrushevk/chat-server-service/internal/repository/model"
	modelService "github.com/vakhrushevk/chat-server-service/internal/service/model"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
)

// ToChatFromRepo - Конвертируем модель репозитория в модель сервиса
func ToChatFromRepo(chat modelRepo.RepoChat) *modelService.ServiceChat {
	var serviceChat modelService.ServiceChat
	serviceChat.ID.Int64 = chat.ID
	serviceChat.Name = chat.Name
	return &serviceChat
}

// FromChatToRepo - TODO: add description
func FromChatToRepo(chat *modelService.ServiceChat) *modelRepo.RepoChat {
	return &modelRepo.RepoChat{Name: chat.Name}
}

// FromServiceMessageToRepoMessage  - TODO: add description
func FromServiceMessageToRepoMessage(message *modelService.ServiceMessage) *modelRepo.RepoMessage {
	repoMessage := &modelRepo.RepoMessage{Sender: message.Sender, Text: message.Text, IDChat: message.IDChat}
	if message.ID.Valid {
		repoMessage.ID = message.ID.Int64
	}

	return repoMessage
}

// FromDescToChat - TODO: add description
func FromDescToChat(request *chat_v1.CreateChatRequest) *modelService.ServiceChat {
	return &modelService.ServiceChat{Name: request.ChatName, UserID: request.IdUsers}
}

// FromDescToServiceMessage - TODO: add description
func FromDescToServiceMessage(request *chat_v1.SendMessageRequest) *modelService.ServiceMessage {
	return &modelService.ServiceMessage{Text: request.GetText(),
		Sender: request.GetFromIdUser(), IDChat: request.GetIdChat()}
}
