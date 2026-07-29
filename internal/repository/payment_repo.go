package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type PaymentRepo interface {
	CreatePayment(m models.Payment) (models.Payment, error)
	GetPaymentById(id int) (models.Payment, error)
	UpdatePayment(id int, status string) error
}

type PaymentRepoImpl struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) *PaymentRepoImpl {
	return &PaymentRepoImpl{db: db}
}

func (r *PaymentRepoImpl) CreatePayment(m models.Payment) (models.Payment, error) {
	err := r.db.Get(&m, `INSERT INTO payments(user_id, subscription_id, amount, status, provider, transaction_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, created_at`, m.UserId, m.SubscriptionId, m.Amount, m.Status, m.Provider, m.TransactionId)
	if err != nil {
		return models.Payment{}, err
	}
	return m, nil
}

func (r *PaymentRepoImpl) GetPaymentById(id int) (models.Payment, error) {
	var m models.Payment
	err := r.db.Get(&m, `SELECT id, user_id, subsctiption_id, amount, status, provider, transaction_id FROM payments WHERE id = $1`, id)
	if err != nil {
		return models.Payment{}, err
	}
	return m, nil
}

func (r *PaymentRepoImpl) UpdatePayment(id int, status string) error {
	_, err := r.db.Exec(`UPDATE payments SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	return nil
}
