package postgres

import (
	"context"
	"errors"
	"github.com/vakhrushevk/local-platform/pkg/db"
	"github.com/vakhrushevk/local-platform/pkg/logger"
	"log"

	"github.com/Masterminds/squirrel"
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
	db db.Client
}

// NewChatRepository - Создаем новый экземлпяр репозитория
func NewChatRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat - Создает чат и заполняет его юзерами? возвращает id чата и ошибку
// TODO: Добавить транзакцию для создания чата и добавления юзеров
func (r *repo) CreateChat(ctx context.Context, chat model.RepoChat) (int64, error) {
	if chat.Name == "" {
		return 0, errors.New("chat name can't be empty")
	}
	query, args, err := squirrel.
		Insert(tableNameChat).
		Columns(nameColumnChat).
		Values(chat.Name).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var id int64
	q := db.Query{Name: "ChatRepository - create", QueryRaw: query}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) AddUserToChat(ctx context.Context, chatID int64, userID int64) error {
	query, args, err := squirrel.Insert(tableNameChatUser).
		Columns(idUserColumnChatUser, idChatColumnChatUser).
		PlaceholderFormat(squirrel.Dollar).Values(userID, chatID).ToSql()

	if err != nil {
		logger.Debug("Ошибка на уровне repository")
		return err
	}

	q := db.Query{Name: "Add User", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
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
	q := db.Query{Name: "ChatRepository - SendMessage", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		log.Println("Failed to save message:", err)
		return err
	}

	return nil
}

func (r *repo) DeleteChat(ctx context.Context, idChat int64) error {
	query, args, err := squirrel.Delete("chat").
		Where(squirrel.Eq{"id": idChat}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}
	q := db.Query{Name: "ChatRepository - DeleteChat", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
