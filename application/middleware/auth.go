package middleware

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
)


var ENV = os.Getenv("ENV")

// nolint
func UserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cookieValue string

		if ENV != "local" {
			cookie, err := c.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			} else {
				cookieValue = cookie.Value
			}
		} else {
			cookieValue := c.QueryParam("cur_user")
			if cookieValue == "" {
				cookieValue = "1"
			}
		}

		userID, err := strconv.Atoi(cookieValue)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("userID", userID)

		return next(c)
	}
}
