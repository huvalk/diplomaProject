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
	"html/template"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"time"
)

type AuthHandler struct {
	useCase auth.UseCase
	tmpl    *Template
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	meta, ok := data.(string)
	if !ok {
		return errors.New("not string as meta")
	}
	return t.templates.Execute(w, template.HTML(meta))
}

func NewAuthHandler(e *echo.Echo, au auth.UseCase) error {
	handler := AuthHandler{
		useCase: au,
		tmpl:    &Template{templates: template.Must(template.ParseFiles("static/index.html"))},
	}

	e.GET("/redirect", handler.RedirectLogin)
	e.GET("/auth", handler.Auth)
	e.GET("/unauth", handler.UnAuth)
	e.GET("/check", handler.Check, middleware.UserID)
	e.GET("/index/*", handler.Static)
	e.Renderer = handler.tmpl
	return nil
}

func (eh *AuthHandler) RedirectLogin(ctx echo.Context) error {
	backTo := url2.QueryEscape(ctx.QueryParam("backTo"))
	url := eh.useCase.MakeAuthUrl(backTo)

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (eh *AuthHandler) Auth(ctx echo.Context) error {
	backTo := ctx.QueryParam("backTo")
	stateTemp := ctx.QueryParam("state")
	code := ctx.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("no code provided"))
	}

	userID, err := eh.useCase.UpdateUserInfo(code, stateTemp, backTo)
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
		Name:     constants.CookieName,
		Value:    t,
		Expires:  time.Now().Add(constants.CookieDuration),
		SameSite: http.SameSiteStrictMode,
		Secure:   globalVars.ENV == constants.PROD,
		HttpOnly: true,
	})
	return ctx.Redirect(http.StatusTemporaryRedirect, globalVars.FRONTEND_URI+backTo)
}

func (eh *AuthHandler) UnAuth(ctx echo.Context) error {
	ctx.SetCookie(&http.Cookie{
		Name:     constants.CookieName,
		Value:    "",
		Expires:  time.Now(),
		SameSite: http.SameSiteStrictMode,
		Secure:   globalVars.ENV == constants.PROD,
		HttpOnly: true,
	})
	return nil
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

func (eh *AuthHandler) Static(ctx echo.Context) error {
	query := ctx.Request().URL.Path
	meta, err := eh.useCase.GenerateMeta(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Render(http.StatusOK, "meta", meta)
}
