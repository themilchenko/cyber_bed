package httpAuth

import (
	"cyber_bed/internal/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) AuthHandler {
	return AuthHandler{
		authUsecase: u,
	}
}

func (h AuthHandler) CreateName(c echo.Context) error {
	name := c.QueryParam("name")
	if err := h.authUsecase.CreateName(name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	responseText := fmt.Sprintf("Hello, %s", name)

	return c.JSON(http.StatusOK, responseText)
}