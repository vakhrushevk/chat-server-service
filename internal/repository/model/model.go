package model

type Chat struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

// TODO: mb stoit udalit
type ChatUser struct {
	Id     int   `db:"id"`
	IdChat int64 `db:"id_chat"`
	IdUser int64 `db:"id_user"`
}

type Message struct {
	Id     int64  `db:"id"`
	Sender int64  `db:"sender"`
	Text   string `db:"text"`
	IdChat int64  `db:"id_chat"`
}
