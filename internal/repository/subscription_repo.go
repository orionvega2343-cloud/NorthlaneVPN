package repository

import (
	"NorthlaneVPN/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type SubscriptionRepo interface {
	CreateSubscription(m models.Subscription) (models.Subscription, error)
	GetBySubscriptionId(id int) (models.Subscription, error)
	GetByUserId(UserId int) (models.Subscription, error)
	UpdateSubscription(m models.Subscription) error
	GetActiveExpiring(before time.Time) ([]models.Subscription, error)
}

type SubscriptionRepoImpl struct {
	db *sqlx.DB
}

func NewSubscriptionRepo(db *sqlx.DB) *SubscriptionRepoImpl {
	return &SubscriptionRepoImpl{db: db}
}

func (r *SubscriptionRepoImpl) CreateSubscription(m models.Subscription) (models.Subscription, error) {
	err := r.db.Get(&m, `INSERT INTO subscription (user_id, tariff_id, server_id, status, starts_at, finishes_at, traffic_used_gb, is_trial, uuid) VALUES($1, $2 ,$3, $4, $5, $6, $7, $8, $9) RETURNING id`, m.UserId, m.TariffId, m.ServerId, m.Status, m.StartsAt, m.FinishesAt, m.TrafficUsedGb, m.IsTrial, m.Uuid)
	if err != nil {
		return models.Subscription{}, err
	}
	return m, nil
}

func (r *SubscriptionRepoImpl) GetBySubscriptionId(id int) (models.Subscription, error) {
	var m models.Subscription
	err := r.db.Get(&m, `SELECT id, user_id, tariff_id, server_id, status, starts_at, finishes_at, traffic_used_gb, is_trial, uuid FROM subscription WHERE id = $1`, id)
	if err != nil {
		return models.Subscription{}, err
	}
	return m, nil
}

func (r *SubscriptionRepoImpl) GetByUserId(UserId int) (models.Subscription, error) {
	var m models.Subscription
	err := r.db.Get(&m, `SELECT id, user_id, tariff_id, server_id, status, starts_at, finishes_at, traffic_used_gb, is_trial, uuid FROM subscription WHERE user_id = $1`, UserId)
	if err != nil {
		return models.Subscription{}, err
	}
	return m, nil
}

func (r *SubscriptionRepoImpl) UpdateSubscription(m models.Subscription) error {
	_, err := r.db.Exec(`UPDATE subscription SET  user_id= $1, tariff_id= $2, server_id= $3, status= $4, starts_at= $5, finishes_at= $6, traffic_used_gb= $7, is_trial= $8, uuid= $9 WHERE id = $10 `, m.UserId, m.TariffId, m.ServerId, m.Status, m.StartsAt, m.FinishesAt, m.TrafficUsedGb, m.IsTrial, m.Uuid, m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepoImpl) GetActiveExpiring(before time.Time) ([]models.Subscription, error) {
	var m []models.Subscription
	err := r.db.Select(&m, `SELECT id, user_id, tariff_id, server_id, status, starts_at, finishes_at,  traffic_used_gb, is_trial, uuid FROM subscription WHERE status IN ('active', 'trial') AND finishes_at < $1`, before)
	if err != nil {
		return nil, err
	}
	return m, nil
}
