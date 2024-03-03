package model

import "time"

type Token struct {
	Id      int       `db:"id"`
	Token   string    `db:"token"`
	Revoked bool      `db:"revoked"`
	Expire  time.Time `db:"expire"`
	UserId  int       `db:"user_id"`
}
