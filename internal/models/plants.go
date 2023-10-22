package models

import (
	"github.com/lib/pq"
)

type UserPlants struct {
	UserID   uint64
	PlantsID pq.Int64Array `gorm:"type:integer[]"`
}

type Plant struct {
	UserID   uint64 `json:"userID"`
	ID       uint64 `json:"id"`
	ImageUrl string `json:"imageUrl"`

	CommonName     string        `json:"common_name"`
	ScientificName []string      `json:"scientific_name"`
	OtherName      []string      `json:"other_name"`
	Cycle          string        `json:"cycle"`
	Watering       string        `json:"watering"`
	Sunlight       []interface{} `json:"sunlight"`
}
