package http

import (
	"diplomaProject/application/event"
	"diplomaProject/application/middleware"
	"diplomaProject/application/models"
	"errors"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
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
	e.GET("/event/:id/users", handler.GetEventUsers)
	e.POST("/event", handler.CreateEvent, middleware.UserID)
	return nil
}

func (eh *EventHandler) GetEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evt, err := eh.useCase.Get(evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) GetEventUsers(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	usrs, err := eh.useCase.GetEventUsers(evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(usrs, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) CreateEvent(ctx echo.Context) error {
	newEvt := &models.Event{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, newEvt); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	//if userID != newEvt.Founder {
	//	fmt.Println("(( ", userID, newEvt.Founder)
	//	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match founder"))
	//}
	newEvt.Founder = userID
	newEvt, err := eh.useCase.Create(newEvt)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(newEvt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
