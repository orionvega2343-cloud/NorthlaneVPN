package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type TariffRepo interface {
	CreateTariff(m models.Tariff) (models.Tariff, error)
	GetAllTariffs() ([]models.Tariff, error)
	GetTariffByID(id int) (models.Tariff, error)
	UpdateTariff(m models.Tariff) error
}

type TariffRepoImpl struct {
	db *sqlx.DB
}

func NewTariffRepo(db *sqlx.DB) *TariffRepoImpl {
	return &TariffRepoImpl{db: db}
}

func (r *TariffRepoImpl) CreateTariff(m models.Tariff) (models.Tariff, error) {
	err := r.db.Get(&m, `INSERT INTO tariffs(name, duration_days, traffic_limit_gb, device_limit, price, is_active) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, m.Name, m.DurationDays, m.TrafficLimitGb, m.DeviceLimit, m.Price, m.IsActive)
	if err != nil {
		return models.Tariff{}, err
	}
	return m, nil
}

func (r *TariffRepoImpl) GetAllTariffs() ([]models.Tariff, error) {
	var m []models.Tariff
	err := r.db.Select(&m, `SELECT id, name, duration_days, traffic_limit_gb, device_limit, price, is_active FROM tariffs WHERE is_active = true`)
	if err != nil {
		return []models.Tariff{}, err
	}
	return m, nil
}

func (r *TariffRepoImpl) GetTariffByID(id int) (models.Tariff, error) {
	var m models.Tariff
	err := r.db.Get(&m, `SELECT id, name, duration_days, traffic_limit_gb, device_limit, price, is_active FROM tariffs WHERE id = $1`, id)
	if err != nil {
		return models.Tariff{}, err
	}
	return m, nil
}

func (r *TariffRepoImpl) UpdateTariff(m models.Tariff) error {
	_, err := r.db.Exec(`UPDATE tariffs SET duration_days= $1, traffic_limit_gb=$2, device_limit=$3, price=$4, is_active=$5 WHERE id =$6`, m.DurationDays, m.TrafficLimitGb, m.DeviceLimit, m.Price, m.IsActive, m.Id)
	if err != nil {
		return err
	}
	return nil
}
