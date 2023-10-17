package httpAuth

import (
	"net/http"
	"time"

	"github.com/cyber_bed/internal/config"
	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase  domain.AuthUsecase
	usersUsecase domain.UsersUsecase

	cookieConfig config.CookieSettings
}

func NewAuthHandler(
	a domain.AuthUsecase,
	u domain.UsersUsecase,
	c config.CookieSettings,
) AuthHandler {
	return AuthHandler{
		authUsecase:  a,
		usersUsecase: u,
		cookieConfig: c,
	}
}

func (h AuthHandler) Auth(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.usersUsecase.GetUserIDBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	if err = h.authUsecase.Auth(cookie.Value); err != nil {
		// If session doesn't exist, create this
		sessionID, err := h.authUsecase.SignUpByID(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		c.SetCookie(h.makeHTTPCookie(sessionID))
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: userID,
	})
}

func (h AuthHandler) SignUp(c echo.Context) error {
	var recievedUser models.User
	if err := c.Bind(&recievedUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if _, err := h.usersUsecase.GetByUsername(recievedUser.Username); err == nil {
	}

	userID, err := h.usersUsecase.CreateUser(recievedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	session, err := h.authUsecase.SignUpByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	c.SetCookie(h.makeHTTPCookie(session))

	return c.JSON(http.StatusOK, models.UserID{
		ID: userID,
	})
}

func (h AuthHandler) Login(c echo.Context) error {
	var authUsr models.AuthUser
	if err := c.Bind(&authUsr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	session, usrID, err := h.authUsecase.Login(authUsr.Username, authUsr.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	c.SetCookie(h.makeHTTPCookie(session))

	return c.JSON(http.StatusOK, models.UserID{
		ID: usrID,
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
		models.DeleteExpire["year"],
		models.DeleteExpire["month"],
		models.DeleteExpire["day"],
	)
	c.SetCookie(h.makeHTTPCookie(cookie.Value))

	return c.JSON(http.StatusOK, []interface{}{})
}
