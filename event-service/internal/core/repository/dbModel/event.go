package dbModel

type PostEvent struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	PostId int    `db:"post_id"`
	UserId int    `db:"user_id"`
}
