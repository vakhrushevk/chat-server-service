package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/repository/model"
)

const (
	tableNameChat        = "chat"
	nameColumnChat       = "name"
	tableNameChatUser    = "chat_user"
	idUserColumnChatUser = "id_user"
	idChatColumnChatUser = "id_chat"

	tableNameMessages    = "messages"
	senderColumnMessages = "sender"
	textColumnMessages   = "text"
	idChatColumnMessages = "id_chat"
)

type repo struct {
	db *pgxpool.Pool
}

// NewChatRepository - Создаем новый экземлпяр репозитория
func NewChatRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat -
func (r *repo) CreateChat(ctx context.Context, chat model.RepoChat, userID []int64) (int64, error) {
	if chat.Name == "" {
		return 0, errors.New("chatservice name cannot be empty")
	}

	query, args, err := squirrel.
		Insert(tableNameChat).
		Columns(nameColumnChat).
		Values(chat.Name).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id").
		ToSql()

	log.Println("Generated query to create a new chatservice:", query, args)

	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Println("Failed to create chatservice:", err)
		return 0, err
	}

	for _, user := range userID {
		query, args, err = squirrel.
			Insert(tableNameChatUser).
			Columns(idUserColumnChatUser, idChatColumnChatUser).
			PlaceholderFormat(squirrel.Dollar).
			Values(user, id).
			ToSql()
		if err != nil {
			return 0, err
		}

		log.Println("Query to add users to chatservice:", query, args)

		_, err = r.db.Exec(ctx, query, args...)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *repo) SendMessage(ctx context.Context, message model.RepoMessage) error {
	query, args, err := squirrel.
		Insert(tableNameMessages).
		Columns(senderColumnMessages, textColumnMessages, idChatColumnMessages).
		Values(message.Sender, message.Text, message.IDChat).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)

	if err != nil {
		log.Println("Failed to save message:", err)
		return err
	}
	return nil
}

func (r repo) DeleteChat(ctx context.Context, idChat int64) error {
	query, args, err := squirrel.Delete("chat").
		Where(squirrel.Eq{"id": idChat}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
