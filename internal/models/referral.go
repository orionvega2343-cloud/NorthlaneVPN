package models

import "time"

type Referral struct {
	Id         int       `db:"id"`
	ReferrerId int       `db:"referrer_id"`
	ReferredId int       `db:"referred_id"`
	Bonus      string    `db:"bonus"`
	CreatedAt  time.Time `db:"created_at"`
}
