package db

import (
	"NorthlaneVPN/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg config.DB) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("name: %s, host, %s, port: %d, username: %s, password: %s, ssl_mode: %s", cfg.Name, cfg.Host, cfg.Port, cfg.User, cfg.Password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
