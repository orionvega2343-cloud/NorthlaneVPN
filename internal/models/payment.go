package models

import "time"

type Payment struct {
	Id             int       `db:"id"`
	UserId         int       `db:"user_id"`
	SubscriptionId int       `db:"subscription_id"`
	Amount         float64   `db:"amount"`
	Status         string    `db:"status"`
	Provider       string    `db:"provider"`
	TransactionId  string    `db:"transaction_id"`
	CreatedAt      time.Time `db:"created_at"`
}
