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

	e.GET("/event/top", handler.GetTopEvents)
	e.GET("/event/:id", handler.GetEvent)
	e.POST("/event/:id/finish", handler.FinishEvent, middleware.UserID)
	e.DELETE("/event/:id/prize", handler.DeletePrize, middleware.UserID)
	e.GET("/event/:id/users", handler.GetEventUsers)
	e.GET("/event/:id/teams", handler.GetEventTeams)
	e.GET("/event/:id/teams/win", handler.GetEventWinnerTeams)
	e.POST("/event", handler.CreateEvent, middleware.UserID)
	e.POST("/event/:id", handler.UpdateEvent, middleware.UserID)
	e.POST("/event/:id/win", handler.SelectWinner, middleware.UserID)
	e.POST("/event/:id/unwin", handler.SelectUnWinner, middleware.UserID)
	e.POST("/event/:id/logo", handler.SetLogo, middleware.UserID)
	e.POST("/event/:id/background", handler.SetBackground, middleware.UserID)
	return nil
}

func (eh *EventHandler) GetTopEvents(ctx echo.Context) error {
	evtArr, err := eh.useCase.GetTopEvents()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evtArr, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
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

func (eh *EventHandler) DeletePrize(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	prArr := &models.PrizeArr{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, prArr); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	evt, err := eh.useCase.RemovePrize(userID, evtID, prArr)
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

func (eh *EventHandler) GetEventTeams(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tms, err := eh.useCase.GetEventTeams(evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(tms, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) GetEventWinnerTeams(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tms, err := eh.useCase.GetEventWinnerTeams(evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(tms, ctx.Response().Writer); err != nil {
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

func (eh *EventHandler) SelectUnWinner(ctx echo.Context) error {
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

	err = eh.useCase.UnSelectWinner(userID, evtID, sel.PrizeID, sel.TeamID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(newEvt, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return ctx.String(200, "OK")
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

func (eh *EventHandler) UpdateEvent(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newEvt := &models.Event{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, newEvt); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newEvt.Id = evtID
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	newEvt, err = eh.useCase.Update(userID, newEvt)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(newEvt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) SetLogo(ctx echo.Context) error {
	eid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	form, _ := ctx.MultipartForm()

	link, err := eh.useCase.SetLogo(userID, eid, form)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(models.Avatar{Avatar: link}, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) SetBackground(ctx echo.Context) error {
	eid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}

	form, _ := ctx.MultipartForm()

	link, err := eh.useCase.SetBackground(userID, eid, form)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(models.Avatar{Avatar: link}, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
