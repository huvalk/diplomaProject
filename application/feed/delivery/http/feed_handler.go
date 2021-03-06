package http

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type FeedHandler struct {
	useCase feed.UseCase
}

func NewFeedHandler(e *echo.Echo, usecase feed.UseCase) error {

	handler := FeedHandler{useCase: usecase}

	e.GET("/feed/:f", handler.GetFeed)
	e.POST("/feed", handler.CreateFeed)
	e.GET("/event/:id/filter", handler.FilterFeed)
	return nil
}

func (fh *FeedHandler) FilterFeed(ctx echo.Context) error {
	evtid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fd, err := fh.useCase.FilterFeed(evtid, ctx.QueryParams())
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(fd, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (fh *FeedHandler) GetFeed(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evt, err := fh.useCase.Get(uid)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(evt, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (fh *FeedHandler) CreateFeed(ctx echo.Context) error {
	add := &models.AddToTeam{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fd, err := fh.useCase.Create(add.UID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(fd, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
