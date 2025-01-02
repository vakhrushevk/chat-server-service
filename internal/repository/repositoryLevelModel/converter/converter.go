package converter

import (
	"github.com/vakhrushevk/chat-server-service/internal/repository/repositoryLevelModel"
	"github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
)

// FromRepositoryLevelToServiceLevel - Конвертируем модель репозитория в модель сервиса
func FromRepositoryLevelToServiceLevel(chat repositoryLevelModel.Chat) serviceLevelModel.Chat {
	return serviceLevelModel.Chat{}
}
