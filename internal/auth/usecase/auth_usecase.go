package authUsecase

import (
	"cyber_bed/internal/domain"
)

type AuthUsecase struct {
	authRepository domain.AuthRepository
}

func NewAuthUsecase(r domain.AuthRepository) AuthUsecase {
	return AuthUsecase{
		authRepository: r,
	}
}

func (u AuthUsecase) CreateName(name string) error {
	if err := u.authRepository.CreateName(name); err != nil {
		return err
	}
	return nil
}
