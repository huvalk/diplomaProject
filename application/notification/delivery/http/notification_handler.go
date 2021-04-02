package http

import (
	"diplomaProject/application/middleware"
	"diplomaProject/application/notification"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

const (
	socketBufferSize = 1024
)

type NotificationHandler struct {
	useCase  notification.UseCase
	upgrader *websocket.Upgrader
}

func NewNotificationHandler(e *echo.Echo, usecase notification.UseCase) error {
	handler := NotificationHandler{
		useCase: usecase,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  socketBufferSize,
			WriteBufferSize: socketBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	e.GET("/notification/:userID", handler.GetPendingNotification, middleware.UserID)
	e.GET("/notification/channel/:userID", handler.ConnectToChannel, middleware.UserID)
	return nil
}

func (eh *NotificationHandler) GetPendingNotification(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	currentUser, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("userID not found"))
	}
	if currentUser != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("not current user"))
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

func (eh *NotificationHandler) ConnectToChannel(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//currentUser, found := ctx.Get("userID").(int)
	//if !found {
	//	log.Println("userID not found")
	//	return echo.NewHTTPError(http.StatusBadRequest, errors.New("userID not found"))
	//}
	//if currentUser != userID {
	//	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("not current user"))
	//}

	ws, err := eh.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.useCase.EnterChannel(userID, ws)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
