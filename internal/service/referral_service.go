package service

import (
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
	"database/sql"
	"errors"
)

type ReferralService interface {
	CreateReferral(m models.Referral) (models.Referral, error)
	GetByReferredId(referredId int) (models.Referral, error)
	UpdateReferral(m models.Referral) error
}

type ReferralServiceImpl struct {
	repo repository.ReferralRepo
}

func NewReferralService(repo repository.ReferralRepo) *ReferralServiceImpl {
	return &ReferralServiceImpl{repo: repo}
}

func (r *ReferralServiceImpl) CreateReferral(m models.Referral) (models.Referral, error) {
	//Проверяем ограничение
	//на приглашенного пользователя
	_, err := r.repo.GetByReferredId(m.ReferredId)
	if errors.Is(err, sql.ErrNoRows) {
		return r.repo.CreateReferral(m)
	}

	if err != nil {
		return models.Referral{}, err
	}

	return m, errors.New("user already referred")
}

func (r *ReferralServiceImpl) GetByReferredId(referredId int) (models.Referral, error) {
	return r.repo.GetByReferredId(referredId)
}

func (r *ReferralServiceImpl) UpdateReferral(m models.Referral) error {
	return r.repo.UpdateReferral(m)
}
