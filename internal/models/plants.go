package models

type UserPlants struct {
	UserID   uint64   `gorm:"primary_key"`
	PlantIDs []uint64 `gorm:"type:integer[]"`
}

type Plant struct {
	UserID     uint64 `json:"userID"`
	ID         uint64 `json:"id"`
	CommonName string `json:"commonName"`
	ImageUrl   string `json:"imageUrl"`
}
