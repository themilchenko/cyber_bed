package httpAuth

import (
	"net/http"
	"time"

	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase  domain.AuthUsecase
	usersUsecase domain.UsersUsecase
}

func NewAuthHandler(a domain.AuthUsecase, u domain.UsersUsecase) AuthHandler {
	return AuthHandler{
		authUsecase:  a,
		usersUsecase: u,
	}
}

func (h AuthHandler) Auth(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if err = h.authUsecase.Auth(cookie.Value); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := h.usersUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: user.ID,
	})
}

// TODO: By creating user need to check if username already exists
func (h AuthHandler) SignUp(c echo.Context) error {
	var recievedUser models.User
	if err := c.Bind(&recievedUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userID, err := h.usersUsecase.CreateUser(recievedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	session, err := h.authUsecase.SignUpByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	c.SetCookie(makeHTTPCookie(session))

	return c.JSON(http.StatusOK, models.UserID{
		ID: userID,
	})
}

func (h AuthHandler) Login(c echo.Context) error {
	var authUsr models.AuthUser
	if err := c.Bind(&authUsr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	session, err := h.authUsecase.Login(authUsr.Username, authUsr.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	c.SetCookie(makeHTTPCookie(session))

	user, err := h.usersUsecase.GetBySessionID(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: user.ID,
	})
}

func (h AuthHandler) Logout(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if err = h.authUsecase.Logout(cookie.Value); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	cookie.Expires = time.Now().AddDate(
		deleteExpire["year"],
		deleteExpire["month"],
		deleteExpire["day"],
	)
	c.SetCookie(makeHTTPCookie(cookie.Value))

	return c.JSON(http.StatusOK, []interface{}{})
}
