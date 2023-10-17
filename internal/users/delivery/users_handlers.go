package httpUsers

import (
	"github.com/cyber_bed/internal/domain"
)

type UsersHandler struct {
	usersUsecase domain.UsersUsecase
}

func NewUsersHandler(u domain.UsersUsecase) UsersHandler {
	return UsersHandler{
		usersUsecase: u,
	}
}
