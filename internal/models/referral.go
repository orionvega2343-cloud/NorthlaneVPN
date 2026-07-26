package models

import "time"

type Referral struct {
	Id         int       `db:"id"`
	ReferrerId int       `db:"referrer_id"`
	Bonus      int       `db:"bonus"`
	CreatedAt  time.Time `db:"created_at"`
}
