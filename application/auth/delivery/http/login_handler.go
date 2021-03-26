package httpAuth

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/middleware"
	"diplomaProject/application/models"
	"diplomaProject/pkg/constants"
	"diplomaProject/pkg/globalVars"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	useCase auth.UseCase
}

func NewAuthHandler(e *echo.Echo, au auth.UseCase) error {
	handler := AuthHandler{
		useCase: au,
	}

	e.GET("/redirect", handler.RedirectLogin)
	e.GET("/auth", handler.Auth)
	e.GET("/check", handler.Check, middleware.UserID)
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

	userID, err := eh.useCase.UpdateUserInfo(code, stateTemp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[constants.UserIdKey] = userID

	t, err := token.SignedString([]byte(globalVars.JWT_SECRET))
	if err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:    constants.CookieName,
		Value:   t,
		Expires: time.Now().Add(constants.CookieDuration),
		//SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})

	return ctx.Redirect(http.StatusTemporaryRedirect, globalVars.FRONTEND_URI)
}

func (eh *AuthHandler) Check(ctx echo.Context) error {
	var user models.AuthUser
	var found bool

	user.Id, found = ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	if _, err := easyjson.MarshalToWriter(user, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
