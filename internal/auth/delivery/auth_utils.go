package httpAuth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
)

const (
	cookieName = "seesion_id"
)

var deleteExpire = map[string]int{
	"year":  0,
	"month": -1,
	"day":   0,
}

func makeHTTPCookie(sessionID string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    sessionID,
		Expires:  time.Now().AddDate(0, 0, 7),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}

func GetCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return cookie, nil
}
