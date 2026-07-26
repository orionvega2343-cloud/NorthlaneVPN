package models

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	TgId      int64     `db:"tg_id"`
	CreatedAt time.Time `db:"created_at"`
	IsBanned  bool      `db:"is_banned"`
}
