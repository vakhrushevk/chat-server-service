package model

// RepoChat - TODO: add description
type RepoChat struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// RepoChatUser - TODO: mb stoit udalit
type RepoChatUser struct {
	ID     int   `db:"id"`
	IDChat int64 `db:"id_chat"`
	IDUser int64 `db:"id_user"`
}

// RepoMessage - TODO: add description
type RepoMessage struct {
	ID     int64  `db:"id"`
	Sender int64  `db:"sender"`
	Text   string `db:"text"`
	IDChat int64  `db:"id_chat"`
}
