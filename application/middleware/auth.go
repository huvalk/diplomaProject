package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
)


var (
	ENV        string = os.Getenv("ENV")
	JWT_SECRET        = os.Getenv("JWT_SECRET")
)

func UserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var userID int

		if ENV == "dev" || ENV == "deploy" {
			cookie, err := c.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				return []byte(JWT_SECRET), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			tokenMap, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("claims jwt error"))
			}

			//TODO переделать
			userIDFloat, ok := tokenMap["userID"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID in jwt token isnt float"))
			}
			tokenMap["userID"] = int(userIDFloat)
			userID, ok = tokenMap["userID"].(int)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID in jwt token isnt float"))
			}
		} else {
			cookieValue := c.QueryParam("cur_user")
			userID, err = strconv.Atoi(cookieValue)
			if err != nil {
				userID = 1
			}
		}

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("userID", userID)

		return next(c)
	}
}
