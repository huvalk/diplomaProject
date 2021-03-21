package httpAuth

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	useCase       auth.UseCase
}

func NewAuthHandler(e *echo.Echo, au auth.UseCase) error {
	handler := AuthHandler{
		useCase: au,
	}

	e.GET("/redirect", handler.RedirectLogin)
	e.GET("/auth", handler.Auth)
	return nil
}

func (eh *AuthHandler) RedirectLogin(ctx echo.Context) error {
	url := eh.useCase.MakeAuthUrl()

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (eh *AuthHandler) Auth(ctx echo.Context) error {
	stateTemp := ctx.QueryParam("state")
	code := ctx.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("no code provided"))
	}

	err := eh.useCase.UpdateUserInfo(code, stateTemp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

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
		Name:     "token",
		Value:    t,
		Expires:  time.Time{},
		MaxAge:   1000000,
		Secure:   false,
		HttpOnly: false,
	})
	return nil
}
