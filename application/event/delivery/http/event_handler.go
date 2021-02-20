package http

import (
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type EventHandler struct {
	useCase event.UseCase
}

func NewEventHandler(e *echo.Echo, usecase event.UseCase) error {

	handler := EventHandler{useCase: usecase}

	e.GET("/event/:id", handler.GetEvent)
	e.POST("/event", handler.CreateEvent)
	return nil
}

func (eh *EventHandler) GetEvent(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evt, err := eh.useCase.Get(uid)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) CreateEvent(ctx echo.Context) error {
	var body []byte
	defer ctx.Request().Body.Close()
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newEvt := &models.Event{}
	err = newEvt.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	err = eh.useCase.Create(newEvt)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(newEvt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}