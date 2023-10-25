package domain

import "github.com/cyber_bed/internal/models"

type PlantsUsecase interface {
	AddPlant(plant models.Plant) error
	GetPlant(userID uint64, plantID int64) (models.Plant, error)
	GetPlants(userID uint64) ([]models.Plant, error)
	DeletePlant(userID, plantID uint64) error
}
