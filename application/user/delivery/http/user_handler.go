package http

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/constants"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
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
	e.GET("/user/:id/events", handler.GetUserEvents)
	e.POST("/login", handler.Login)
	e.PUT("/user/:id", handler.Update)
	e.POST("/user/:id/image", handler.SetImage)
	e.POST("/event/:evtid/join", handler.JoinEvent)
	e.POST("/event/:evtid/leave", handler.LeaveEvent)
	return nil
}

func (uh *UserHandler) GetUserEvents(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evtArr, err := uh.useCase.GetUserEvents(uid)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evtArr, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (uh *UserHandler) SetImage(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	form, _ := ctx.MultipartForm()

	err = uh.useCase.SetImage(uid, form)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return echo.NewHTTPError(http.StatusOK, "OK")
}

func (uh *UserHandler) JoinEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	add := &models.AddToTeam{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = uh.useCase.JoinEvent(add.UID, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return ctx.String(200, "OK")
}

func (uh *UserHandler) LeaveEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	add := &models.AddToTeam{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = uh.useCase.LeaveEvent(add.UID, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return ctx.String(200, "OK")
}

func (uh *UserHandler) Login(ctx echo.Context) error {
	usr := &models.User{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, usr); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sessionId, token, err := uh.useCase.Login(usr.FirstName, usr.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	uh.setCookie(ctx, sessionId)
	uh.setCsrfToken(ctx, token)

	return ctx.String(200, "OK")
}

func (uh *UserHandler) Update(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	usr := &models.User{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, usr); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	usr.Id = uid
	usr, err = uh.useCase.Update(usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(usr, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (uh *UserHandler) Profile(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	usr, err := uh.useCase.GetForFeed(uid)
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
