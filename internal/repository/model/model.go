package model

// Chat - TODO: add description
type Chat struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// ChatUser - TODO: mb stoit udalit
type ChatUser struct {
	ID     int   `db:"id"`
	IDChat int64 `db:"id_chat"`
	IDUser int64 `db:"id_user"`
}

// Message - TODO: add description
type Message struct {
	ID     int64  `db:"id"`
	Sender int64  `db:"sender"`
	Text   string `db:"text"`
	IDChat int64  `db:"id_chat"`
}
