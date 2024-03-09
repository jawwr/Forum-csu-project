package dbModel

type Post struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Text   string `db:"text"`
	UserId int    `db:"user_id"`
}
