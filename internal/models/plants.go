package models

import (
	"github.com/lib/pq"
)

type UserPlants struct {
	UserID   uint64
	PlantsID pq.Int64Array `gorm:"type:integer[]"`
}

type Plant struct {
	UserID     uint64 `json:"userID"`
	ID         uint64 `json:"id"`
	CommonName string `json:"commonName"`
	ImageUrl   string `json:"imageUrl"`
}
