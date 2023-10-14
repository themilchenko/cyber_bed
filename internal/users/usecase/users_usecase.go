package usersUsecase

import (
	"github.com/cyber_bed/internal/crypto"
	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
)

type UsersUsecase struct {
	usersRepository domain.UsersRepository
}

func NewUsersUsecase(r domain.UsersRepository) domain.UsersUsecase {
	return UsersUsecase{
		usersRepository: r,
	}
}

func (u UsersUsecase) CreateUser(user models.User) (uint64, error) {
	hash, err := crypto.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hash

	id, err := u.usersRepository.Create(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u UsersUsecase) GetBySessionID(sessionID string) (models.User, error) {
	user, err := u.usersRepository.GetBySessionID(sessionID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u UsersUsecase) GetByUsername(username string) (models.User, error) {
	user, err := u.usersRepository.GetByUsername(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u UsersUsecase) GetByID(userID uint64) (models.User, error) {
	user, err := u.usersRepository.GetByID(userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
