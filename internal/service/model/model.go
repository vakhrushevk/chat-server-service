package model

import "database/sql"

// ServiceChat - TODO: add description
type ServiceChat struct {
	ID     sql.NullInt64
	Name   string
	UserID []int64
}

// ServiceChatUser - TODO: mb stoit udalit
type ServiceChatUser struct {
	ID     sql.NullInt64
	IDChat int64
	IDUser int64
}

// ServiceMessage - TODO: add description
type ServiceMessage struct {
	ID     sql.NullInt64
	Sender int64  `db:"sender"`
	Text   string `db:"text"`
	IDChat int64  `db:"id_chat"`
}
