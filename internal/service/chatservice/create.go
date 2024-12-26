package chatservice

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/vakhrushevk/chat-server-service/internal/service/model"
	"github.com/vakhrushevk/chat-server-service/internal/service/model/converter"
)

// CreateChat - TODO: Add description
func (s *serv) CreateChat(ctx context.Context, chat *model.ServiceChat) (int64, error) {
	// TODO: Обработка ошибок из репозитория
	log.Println("Я тут был")
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repositoy.CreateChat(ctx, *converter.FromChatToRepo(chat), chat.UserID)
		if errTx != nil {
			return errTx
		}
		fmt.Println("[DEBUG] ID: ", id)
		return errors.New("test error")
		id, errTx = s.repositoy.CreateChat(ctx, *converter.FromChatToRepo(chat), chat.UserID)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		fmt.Println("[DEBUG] Error: ", err)
		return 0, err
	}
	return id, nil
}
