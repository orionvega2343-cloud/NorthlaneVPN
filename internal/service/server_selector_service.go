package service

import (
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
	"fmt"
	"net"
	"time"
)

type ServerService interface {
	CreateServer(m models.Server) (models.Server, error)
	GetAllServers() ([]models.Server, error)
	UpdateServer(m models.Server) error
}

type ServerServiceImpl struct {
	repo repository.ServerRepo
}

func NewServerService(repo repository.ServerRepo) *ServerServiceImpl {
	return &ServerServiceImpl{repo: repo}
}

func (s *ServerServiceImpl) CreateServer(m models.Server) (models.Server, error) {
	str := fmt.Sprintf("%s:%d", m.Host, m.Port)
	conn, err := net.DialTimeout("tcp", str, 10*time.Second)
	if err != nil {
		m.Status = "banned"
	} else {
		defer conn.Close()
		m.Status = "active"
	}

	server, err := s.repo.CreateServer(m)
	if err != nil {
		return m, err
	}
	return server, nil
}

func (s *ServerServiceImpl) GetAllServers() ([]models.Server, error) {
	return s.repo.GetAllServers()
}

func (s *ServerServiceImpl) UpdateServer(m models.Server) error {
	return s.repo.UpdateServer(m)
}
