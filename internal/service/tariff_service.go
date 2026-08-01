package service

import (
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
)

// Tariff - почти неизменяемая сущность приложения,
// используется как начинка под subscription
type TariffService interface {
	CreateTariff(m models.Tariff) (models.Tariff, error)
	GetAllTariffs() ([]models.Tariff, error)
	UpdateTariff(m models.Tariff) error
}

type TariffServiceImpl struct {
	repo repository.TariffRepo
}

func NewTariffService(repo repository.TariffRepo) *TariffServiceImpl {
	return &TariffServiceImpl{repo: repo}
}

func (t *TariffServiceImpl) CreateTariff(m models.Tariff) (models.Tariff, error) {
	return t.repo.CreateTariff(m)
}

func (t *TariffServiceImpl) GetAllTariffs() ([]models.Tariff, error) {
	return t.repo.GetAllTariffs()
}

func (t *TariffServiceImpl) UpdateTariff(m models.Tariff) error {
	return t.repo.UpdateTariff(m)
}
