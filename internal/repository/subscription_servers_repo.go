package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type SubscriptionServersRepo interface {
	GetServersForSubscription(subscriptionId int) ([]models.Server, error)
}

type SubscriptionServersRepoImpl struct {
	db *sqlx.DB
}

func NewSubscriptionServersRepo(db *sqlx.DB) *SubscriptionServersRepoImpl {
	return &SubscriptionServersRepoImpl{db: db}
}

func (r *SubscriptionServersRepoImpl) GetServersForSubscription(subscriptionId int) ([]models.Server, error) {
	var servers []models.Server
	err := r.db.Select(&servers, `SELECT id, host, port, protocol, name, reality_key, sni FROM servers JOIN subscription_servers ON servers.id = subscription_servers.server_id WHERE subscription_servers.subscription_id = $1`, subscriptionId)
	if err != nil {
		return nil, err
	}
	return servers, nil
}
