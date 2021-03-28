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
	e.GET("/event/:id/finish", handler.FinishEvent, middleware.UserID)
	e.GET("/event/:id/users", handler.GetEventUsers)
	e.POST("/event", handler.CreateEvent, middleware.UserID)
	e.POST("/event/:id/win", handler.SelectWinner, middleware.UserID)
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

func (eh *EventHandler) FinishEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	evt, err := eh.useCase.Finish(userID, evtID)
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

func (eh *EventHandler) SelectWinner(ctx echo.Context) error {
	sel := &models.SelectWinner{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, sel); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.useCase.SelectWinner(userID, evtID, sel.PrizeID, sel.TeamID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(newEvt, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return ctx.String(200, "OK")
}
