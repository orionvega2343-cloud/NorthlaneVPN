package models

import "time"

type Server struct {
	Id         int       `db:"id"`
	Host       string    `db:"host"`
	Port       int       `db:"port"`
	Protocol   string    `db:"protocol"`
	Status     string    `db:"status"`
	LoadScore  int       `db:"load_score"`
	Region     string    `db:"region"`
	CreatedAt  time.Time `db:"created_at"`
	Name       string    `db:"name"`
	RealityKey string    `db:"reality_key"`
	Sni        string    `db:"sni"`
}
