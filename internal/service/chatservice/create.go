package chatservice

import (
	"context"
	"fmt"

	"github.com/vakhrushevk/chat-server-service/internal/service/model"
	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
)

// CreateChat - Create a chat instance from the given configuration
func (s *serv) CreateChat(ctx context.Context, chat *model.ServiceChat) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repository.CreateChat(ctx, *converter.FromChatToRepo(chat))
		if errTx != nil {
			return errTx
		}

		for _, UID := range chat.UserID {
			errTx = s.repository.AddUserToChat(ctx, id, UID)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("[DEBUG] Error: ", err)
		//	logger.Error("Error serviceLevel: ", sl.Err(err))
		return 0, err
	}

	return id, nil
}
