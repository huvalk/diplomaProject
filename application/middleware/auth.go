package middleware

import (
	"github.com/labstack/echo"
	"strconv"
)

func UserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int
		cookie, err := c.Cookie("token")
		if err != nil {
			userID = 1
		} else {
			userID, err = strconv.Atoi(cookie.Value)
			if err != nil {
				userID = 1
			}
		}

		c.Set("userID", userID)

		return next(c)
	}
}
