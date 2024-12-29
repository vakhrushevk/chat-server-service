package chatservice

import "context"

// DeleteChat - it's DELETE CHAT!!!!
func (s *serv) DeleteChat(ctx context.Context, idChat int64) error {
	return s.repository.DeleteChat(ctx, idChat)
}
