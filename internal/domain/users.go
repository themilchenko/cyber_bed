package domain

import "github.com/cyber_bed/internal/models"

type UsersUsecase interface {
	CreateUser(user models.User) (uint64, error)

	GetUserIDBySessionID(sessionID string) (uint64, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetByID(userID uint64) (models.User, error)
}
