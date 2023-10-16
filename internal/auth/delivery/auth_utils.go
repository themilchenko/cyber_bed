package httpAuth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	cookieName = "session_id"
)

func (h AuthHandler) makeHTTPCookie(sessionID string) *http.Cookie {
	return &http.Cookie{
		Name:  cookieName,
		Value: sessionID,
		Expires: time.Now().
			AddDate(int(h.config.CookieSettings.ExpireDate.Years),
				int(h.config.CookieSettings.ExpireDate.Months),
				int(h.config.CookieSettings.ExpireDate.Days)),
		Secure:   h.config.CookieSettings.Secure,
		HttpOnly: h.config.CookieSettings.HttpOnly,
		SameSite: http.SameSiteNoneMode,
	}
}

func GetCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}
