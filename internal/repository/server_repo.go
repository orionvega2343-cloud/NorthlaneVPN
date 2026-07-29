package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type ServerRepo interface {
	CreateServer(m models.Server) (models.Server, error)
	GetAllServers() ([]models.Server, error)
	UpdateServer(m models.Server) error
}

type ServerRepoImpl struct {
	db *sqlx.DB
}

func NewServerRepo(db *sqlx.DB) *ServerRepoImpl {
	return &ServerRepoImpl{db: db}
}

func (r *ServerRepoImpl) CreateServer(m models.Server) (models.Server, error) {
	err := r.db.Get(&m, `INSERT INTO servers(host, port, protocol, status, load_score, region) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, created_at`, m.Host, m.Port, m.Protocol, m.Status, m.LoadScore, m.Region)
	if err != nil {
		return models.Server{}, err
	}
	return m, nil
}

func (r *ServerRepoImpl) GetAllServers() ([]models.Server, error) {
	var servers []models.Server
	err := r.db.Select(&servers, `SELECT id, host, port, protocol, status, load_score, region, created_at FROM servers WHERE status IN ('active')`)
	if err != nil {
		return servers, err
	}
	return servers, nil
}

func (r *ServerRepoImpl) UpdateServer(m models.Server) error {
	_, err := r.db.Exec(`UPDATE servers SET host = $1, port = $2, protocol = $3, status = $4, load_score = $5 WHERE id = $6`, m.Host, m.Port, m.Protocol, m.Status, m.LoadScore, m.Id)
	if err != nil {
		return err
	}
	return nil
}
