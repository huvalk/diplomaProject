package http

import (
	"diplomaProject/application/notification"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type EventHandler struct {
	useCase notification.UseCase
	upgrader *websocket.Upgrader
}

const (
	socketBufferSize = 1024
)

func NewNotificationHandler(e *echo.Echo, usecase notification.UseCase) error {
	handler := EventHandler{
		useCase: usecase,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  socketBufferSize,
			WriteBufferSize: socketBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	e.GET("/notification/:userID", handler.GetPendingNotification)
	e.POST("/notification/channel/:userID", handler.ConnectToChannel)
	return nil
}

func (eh *EventHandler) GetPendingNotification(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	notifications, err := eh.useCase.GetPendingNotification(userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(notifications, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (eh *EventHandler) ConnectToChannel(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ws, err := eh.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.useCase.EnterChannel(userId, ws)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
