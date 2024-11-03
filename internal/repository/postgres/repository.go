package postgres

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/repository/model"
	"log"
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

// NewRepository - Создаем новый экземлпяр репозитория
func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat -
func (r *repo) CreateChat(ctx context.Context, chat model.Chat, userId []int64) (int64, error) {
	if chat.Name == "" {
		return 0, errors.New("chat name cannot be empty")
	}

	query, args, err := squirrel.
		Insert(tableNameChat).
		Columns(nameColumnChat).
		Values(chat.Name).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id").
		ToSql()

	log.Println("Generated query to create a new chat:", query, args)

	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Println("Failed to create chat:", err)
		return 0, err
	}

	for _, user := range userId {
		query, args, err := squirrel.
			Insert(tableNameChatUser).
			Columns(idUserColumnChatUser, idChatColumnChatUser).
			PlaceholderFormat(squirrel.Dollar).
			Values(user, id).
			ToSql()

		log.Println("Query to add users to chat:", query, args)

		_, err = r.db.Exec(ctx, query, args...)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *repo) SendMessage(ctx context.Context, message model.Message) error {
	query, args, err := squirrel.
		Insert(tableNameMessages).
		Columns(senderColumnMessages, textColumnMessages, idChatColumnMessages).
		Values(message.Sender, message.Text, message.IdChat).
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
