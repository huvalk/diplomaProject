package http

import (
	"diplomaProject/application/middleware"
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/constants"
	"diplomaProject/pkg/crypto"
	"errors"
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
	e.GET("/founder/events", handler.GetCreatedEvents, middleware.UserID)
	e.GET("/user/:id/events", handler.GetUserEvents)
	e.GET("/event/:eventID/user/search", handler.FindUserByTag)
	// Логин реализован в auth
	//e.POST("/login", handler.Login)
	e.PUT("/user/:id", handler.Update, middleware.UserID)
	e.POST("/user/:id/image", handler.SetImage, middleware.UserID)
	e.POST("/event/:evtid/join", handler.JoinEvent, middleware.UserID)
	e.POST("/event/:evtid/leave", handler.LeaveEvent, middleware.UserID)
	return nil
}

func (uh *UserHandler) GetCreatedEvents(ctx echo.Context) error {
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	evtArr, err := uh.useCase.GetFounderEvents(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evtArr, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
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

func (uh *UserHandler) FindUserByTag(ctx echo.Context) error {
	eid, err := strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tag := ctx.QueryParam("tag")
	//if err != nil {
	//	log.Println(err)
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	users, err := uh.useCase.SearchUserByTag(eid, tag)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(users, ctx.Response().Writer); err != nil {
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
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	if userID != uid {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	}

	form, _ := ctx.MultipartForm()

	link, err := uh.useCase.SetImage(uid, form)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(models.Avatar{Avatar: link}, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (uh *UserHandler) JoinEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	evt, err := uh.useCase.GetBDEvent(evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if evt.IsPrivate {
		if !crypto.CheckToken(strconv.Itoa(evtID), ctx.QueryParam("secret")) {
			return echo.NewHTTPError(http.StatusUnauthorized, errors.New("not valid secret"))
		}
	}

	err = uh.useCase.JoinEvent(userID, evtID)
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
	//add := &models.AddToTeam{}
	//if err := easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
	//	log.Println(err)
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	//if userID != add.UID {
	//	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	//}

	err = uh.useCase.LeaveEvent(userID, evtID)
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
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	if userID != uid {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	}

	usr := &models.User{
		WorkPlace:   "SAVE_THIS_FIELD",
		Description: "SAVE_THIS_FIELD",
		Bio:         "SAVE_THIS_FIELD",
		Vk:          "SAVE_THIS_FIELD",
		Tg:          "SAVE_THIS_FIELD",
		Git:         "SAVE_THIS_FIELD",
	}
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
