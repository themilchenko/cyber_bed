package authUsecase

import (
	"errors"
	"time"

	"github.com/cyber_bed/internal/crypto"
	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	authRepository   domain.AuthRepository
	usersRepoisitory domain.UsersRepository
}

func NewAuthUsecase(r domain.AuthRepository, u domain.UsersRepository) domain.AuthUsecase {
	return AuthUsecase{
		authRepository:   r,
		usersRepoisitory: u,
	}
}

func (u AuthUsecase) generateCookie(userID uint64) models.Cookie {
	return models.Cookie{
		UserID:     userID,
		Value:      uuid.New().String(),
		ExpireDate: time.Now().AddDate(0, 0, 7),
	}
}

func (u AuthUsecase) Auth(sessionID string) error {
	if _, err := u.usersRepoisitory.GetBySessionID(sessionID); err != nil {
		return err
	}
	return nil
}

func (u AuthUsecase) SignUpByID(userID uint64) (string, error) {
	session, err := u.authRepository.CreateSession(u.generateCookie(userID))
	if err != nil {
		return "", err
	}
	return session, nil
}

func (u AuthUsecase) Login(login, password string) (string, error) {
	user, err := u.usersRepoisitory.GetByUsername(login)
	if err != nil {
		return "", err
	}

	if !crypto.CheckHash(user.Password, password) {
		// TODO: Add custom errors
		return "", errors.New("wrong passord")
	}

	session, err := u.authRepository.CreateSession(u.generateCookie(user.ID))
	if err != nil {
		return "", err
	}
	return session, nil
}

func (u AuthUsecase) Logout(sessionID string) error {
	if err := u.authRepository.DeleteBySessionID(sessionID); err != nil {
		return err
	}
	return nil
}
