package models

import "time"

type Subscription struct {
	Id            int       `db:"id"`
	UserId        int       `db:"user_id"`
	TariffId      int       `db:"tariff_id"`
	ServerId      int       `db:"server_id"`
	Status        string    `db:"status"`
	StartsAt      time.Time `db:"starts_at"`
	FinishesAt    time.Time `db:"finishes_at"`
	TrafficUsedGb int       `db:"traffic_used_gb"`
	IsTrial       bool      `db:"is_trial"`
	Uuid          string    `db:"uuid"`
}
