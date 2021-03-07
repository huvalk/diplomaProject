package http

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type InviteHandler struct {
	invite       invite.UseCase
	notification notification.UseCase
}

func NewInviteHandler(e *echo.Echo, iu invite.UseCase, nu notification.UseCase) error {
	handler := InviteHandler{
		invite: iu,
		notification: nu,
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

	inv.OwnerID, err = 1, nil
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	notify, err := eh.invite.InviteUser(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if notify {
		err = eh.notification.SendInviteNotificationTo(inv.GuestID, "Оповещение о приглашении")
		if err != nil {
			log.Println("Notification wasnt sent: ", err)
		}
	}

	return nil
}

func (eh *InviteHandler) InviteTeam(ctx echo.Context) (err error) {
	inv := &models.Invitation{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, inv); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inv.OwnerID, err = 1, nil
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	notify, err := eh.invite.InviteTeam(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if notify {
		err = eh.notification.SendInviteNotificationTo(inv.GuestID, "Оповещение о приглашении")
		if err != nil {
			log.Println("Notification wasnt sent: ", err)
		}
	}

	return nil
}

func (eh *InviteHandler) GetInvitedUser(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	inv.OwnerID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := eh.invite.GetInvitedUser(inv)
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

	inv.OwnerID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	teams, err := eh.invite.GetInvitedTeam(inv)
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

	inv.OwnerID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := eh.invite.GetInvitationUser(inv)
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

	inv.OwnerID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	teams, err := eh.invite.GetInvitationTeam(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(teams, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}