package httpAuth

import (
	"diplomaProject/application/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
}

func NewAuthHandler(e *echo.Echo) error {
	handler := AuthHandler{
	}

	e.POST("/login", handler.Login)
	return nil
}

func (eh *AuthHandler) Login(ctx echo.Context) error {
	user := &models.AuthUser{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:       "token",
		Value:      t,
		Expires:    time.Time{},
		MaxAge:     1000000,
		Secure:     false,
		HttpOnly:   false,
	})
	return nil
}
