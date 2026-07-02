package domain

import "time"

type Club struct {
	ID           int64
	Name         string
	City         string
	Address      string
	Description  string
	WorkingHours string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
