package http

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/constants"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	useCase user.UseCase
}

func NewUserHandler(e *echo.Echo, usecase user.UseCase) error {

	handler := UserHandler{useCase: usecase}

	e.GET("/user/:id", handler.Profile)
	e.POST("/login", handler.Login)
	e.POST("/event/:id/join", handler.JoinEvent)
	return nil
}

func (uh *UserHandler) JoinEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var body []byte
	defer ctx.Request().Body.Close()
	body, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	add := &models.AddToTeam{}
	err = add.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	err = uh.useCase.JoinEvent(add.UID, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(usr, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return ctx.String(200, "OK")
}

func (uh *UserHandler) Login(ctx echo.Context) error {
	usr := &models.VkUser{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, usr); err != nil {
		log.Println(err)
		return ctx.String(499, err.Error())
	}
	sessionId, token, err := uh.useCase.Login(usr.FirstName, usr.Email)
	if err != nil {
		return ctx.String(498, err.Error())
	}

	uh.setCookie(ctx, sessionId)
	uh.setCsrfToken(ctx, token)

	return ctx.String(200, "OK")
}

func (uh *UserHandler) Profile(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	usr, err := uh.useCase.Get(uid)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(usr, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (uh *UserHandler) setCookie(ctx echo.Context, sessionId string) {
	cookie := &http.Cookie{
		Name:    constants.CookieName,
		Value:   sessionId,
		Path:    "/",
		Expires: time.Now().Add(constants.CookieDuration),
		//SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}
	ctx.SetCookie(cookie)
}

func (uh *UserHandler) setCsrfToken(ctx echo.Context, token string) {
	cookie := &http.Cookie{
		Name:    constants.CSRFHeader,
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Hour),
		//SameSite: http.SameSiteStrictMode,
	}
	ctx.SetCookie(cookie)
}
