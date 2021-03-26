package middleware

import (
	"diplomaProject/pkg/constants"
	"diplomaProject/pkg/globalVars"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func UserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var userID int

		if globalVars.ENV == "dev" || globalVars.ENV == "deploy" {
			cookie, err := c.Cookie(constants.CookieName)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				return []byte(globalVars.JWT_SECRET), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			tokenMap, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("claims jwt error"))
			}

			//TODO переделать
			userIDFloat, ok := tokenMap[constants.UserIdKey].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID in jwt token isnt float"))
			}
			tokenMap[constants.UserIdKey] = int(userIDFloat)
			userID, ok = tokenMap[constants.UserIdKey].(int)
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

		c.Set(constants.UserIdKey, userID)

		return next(c)
	}
}
