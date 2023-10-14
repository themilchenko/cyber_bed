package domain

import (
	"github.com/cyber_bed/internal/models"
)

type AuthRepository interface {
	CreateSession(cookie models.Cookie) (string, error)
	DeleteBySessionID(sessionID string) error
}

type UsersRepository interface {
	Create(user models.User) (uint64, error)

	GetByUsername(username string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
}

type PlantsRepository interface {
	CreateUserPlantsRelations(userID uint64, plantID []uint64) error
	AddUserPlantsRelations(userID uint64, plantsID []uint64) error
	GetPlantsByID(userID uint64) ([]uint64, error)
}
