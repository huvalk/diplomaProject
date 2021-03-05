package http

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type InviteHandler struct {
	useCase invite.UseCase
	upgrader *websocket.Upgrader
}

func NewInviteHandler(e *echo.Echo, usecase invite.UseCase) error {
	handler := InviteHandler{
		useCase: usecase,
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	e.POST("/event/:eventID/user/:userID/invite", handler.InviteUser)
	e.POST("/event/:eventID/team/:teamID/invite", handler.InviteTeam)
	e.GET("/event/:eventID/invited/user", handler.GetInvitedUser)
	e.GET("/event/:eventID/invited/team", handler.GetInvitedTeam)
	e.GET("/event/:eventID/invitation/user", handler.GetInvitationUser)
	e.GET("/event/:eventID/invitation/team", handler.GetInvitationTeam)
	return nil
}

func (eh *InviteHandler) InviteUser(ctx echo.Context) (err error) {
	inv := &models.Invitation{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, inv); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inv.UserID, err = 1, nil
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.useCase.InviteUser(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (eh *InviteHandler) InviteTeam(ctx echo.Context) (err error) {
	inv := &models.Invitation{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, inv); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inv.UserID, err = 1, nil
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.useCase.InviteTeam(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (eh *InviteHandler) GetInvitedUser(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	inv.UserID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := eh.useCase.GetInvitedUser(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(users, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (eh *InviteHandler) GetInvitedTeam(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	inv.UserID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	teams, err := eh.useCase.GetInvitedTeam(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(teams, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (eh *InviteHandler) GetInvitationUser(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	inv.UserID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := eh.useCase.GetInvitationUser(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(users, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (eh *InviteHandler) GetInvitationTeam(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	inv.UserID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	teams, err := eh.useCase.GetInvitationTeam(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(teams, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}