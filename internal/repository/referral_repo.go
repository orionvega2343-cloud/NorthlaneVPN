package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type ReferralRepo interface {
	CreateReferral(m models.Referral) (models.Referral, error)
	GetByReferredId(referredId int) (models.Referral, error)
	UpdateReferral(m models.Referral) error
}

type ReferralRepoImpl struct {
	db *sqlx.DB
}

func NewReferralRepo(db *sqlx.DB) *ReferralRepoImpl {
	return &ReferralRepoImpl{db: db}
}

func (r *ReferralRepoImpl) CreateReferral(m models.Referral) (models.Referral, error) {
	err := r.db.Get(&m, `INSERT INTO referrals(referrer_id, referred_id, bonus) VALUES($1, $2, $3) RETURNING id,  created_at`, m.ReferrerId, m.ReferredId, m.Bonus)
	if err != nil {
		return models.Referral{}, err
	}
	return m, nil
}

func (r *ReferralRepoImpl) GetByReferredId(referredId int) (models.Referral, error) {
	var m models.Referral
	err := r.db.Get(&m, `SELECT id, referrer_id, referred_id, bonus, created_at FROM referrals WHERE referred_id = $1`, referredId)
	if err != nil {
		return models.Referral{}, err
	}
	return m, nil
}

func (r *ReferralRepoImpl) UpdateReferral(m models.Referral) error {
	_, err := r.db.Exec(`UPDATE referrals SET  referrer_id = $1, referred_id =$2, bonus = $3  WHERE id = $4`, m.ReferrerId, m.ReferredId, m.Bonus, m.Id)
	if err != nil {
		return err
	}
	return nil
}
