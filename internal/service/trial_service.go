package service

import (
	"NorthlaneVPN/internal/models"
	"database/sql"
	"errors"
	"time"
)

// TODO: вынести id тарифа триала в конфиг
const trialTariffId = 1

// CreateTrial - обертка над CreateSubscription,
// доступна пользователю только один раз
func (s *SubscriptionServiceImpl) CreateTrial(userId int) (models.Subscription, error) {
	//Проверяем, что у пользователя еще нет подписки -
	//триал выдается только один раз
	_, err := s.subRepo.GetByUserId(userId)
	if err == nil {
		return models.Subscription{}, errors.New("trial already used")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return models.Subscription{}, err
	}

	return s.CreateSubscription(models.Subscription{
		UserId:   userId,
		TariffId: trialTariffId,
		StartsAt: time.Now(),
		IsTrial:  true,
	})
}
