package models

type Tariff struct {
	Id             int     `db:"id"`
	Name           string  `db:"name"`
	DurationDays   int     `db:"duration_days"`
	TrafficLimitGb int     `db:"traffic_limit_gb"`
	DeviceLimit    int     `db:"device_limit"`
	Price          float64 `db:"price"`
	IsActive       bool    `db:"is_active"`
}
