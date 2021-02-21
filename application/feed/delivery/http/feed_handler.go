package http

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type FeedHandler struct {
	useCase feed.UseCase
}

func NewFeedHandler(e *echo.Echo, usecase feed.UseCase) error {

	handler := FeedHandler{useCase: usecase}

	e.GET("/feed/:id", handler.GetFeed)
	e.POST("/feed", handler.CreateFeed)
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
	var body []byte
	defer ctx.Request().Body.Close()
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newFeed := &models.AddToTeam{}
	err = newFeed.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	fd, err := fh.useCase.Create(newFeed.UID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(fd, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
