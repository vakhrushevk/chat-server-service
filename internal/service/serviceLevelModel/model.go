package serviceLevelModel

import (
	"time"
)

type Chat struct {
	ID        int64
	ChatInfo  ChatInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChatInfo struct {
	Name           string
	CreatedBy      int64
	ChatMembersIds []int64
}

type ChatMember struct {
	ID             int64
	ChatMemberInfo ChatMemberInfo
	JoinedAt       time.Time
}

type ChatMemberInfo struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}
