package dbModel

type Post struct {
	Id     int    `db:"id"`
	Text   string `db:"text"`
	UserId int    `db:"user_id"`
}
