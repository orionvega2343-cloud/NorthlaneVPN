package service

import (
	"NorthlaneVPN/internal/config"
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
	"time"
)

type SubscriptionService interface {
	CreateSubscription(m models.Subscription) (models.Subscription, error)
	CreateTrial(userId int) (models.Subscription, error)
	GetBySubscriptionId(id int) (models.Subscription, error)
	GetByUserId(UserId int) (models.Subscription, error)
	UpdateSubscription(m models.Subscription) error
	GetActiveExpiring() ([]models.Subscription, error)
}

type SubscriptionServiceImpl struct {
	tRepo   repository.TariffRepo
	subRepo repository.SubscriptionRepo
	cfg     *config.Config
}

func NewSubscriptionServiceImpl(tRepo repository.TariffRepo, subRepo repository.SubscriptionRepo) *SubscriptionServiceImpl {
	return &SubscriptionServiceImpl{tRepo: tRepo, subRepo: subRepo}
}

func (s *SubscriptionServiceImpl) CreateSubscription(m models.Subscription) (models.Subscription, error) {
	//Получаем id тарифа, чтобы понять,
	//какой тип подписки выбрал пользователь
	t, err := s.tRepo.GetTariffByID(m.TariffId)
	if err != nil {
		return models.Subscription{}, err
	}

	//Считаем срок подписки
	m.FinishesAt = m.StartsAt.Add(t.DurationDays)

	//TODO: привязать реальный id сервера
	m.ServerId = 1

	if m.IsTrial {
		m.Status = "trial"
	} else {
		m.Status = "active"
	}

	//Создаем подписку
	sub, err := s.subRepo.CreateSubscription(m)
	if err != nil {
		return models.Subscription{}, err
	}
	return sub, nil
}

func (s *SubscriptionServiceImpl) GetBySubscriptionId(id int) (models.Subscription, error) {
	return s.subRepo.GetBySubscriptionId(id)
}

func (s *SubscriptionServiceImpl) GetByUserId(userId int) (models.Subscription, error) {
	return s.subRepo.GetByUserId(userId)
}

func (s *SubscriptionServiceImpl) UpdateSubscription(m models.Subscription) error {
	return s.subRepo.UpdateSubscription(m)
}

func (s *SubscriptionServiceImpl) GetActiveExpiring() ([]models.Subscription, error) {
	//Парсим конфиг в Go syntax
	timeout := time.Now().Add(s.cfg.Notifier.ExpiresAt)

	e, err := s.subRepo.GetActiveExpiring(timeout)
	if err != nil {
		return nil, err
	}
	return e, nil
}
