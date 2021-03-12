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

	e.POST("/event/:eventID/user/:userID/invite", handler.Invite)
	e.POST("/event/:eventID/user/:userID/uninvite", handler.UnInvite)
	e.POST("/event/:eventID/user/:userID/deny", handler.Deny)
	e.GET("/event/:eventID/invited/user/:userID", handler.IsInvited)
	e.GET("/event/:eventID/invited/users", handler.GetInvitedUser)
	e.GET("/event/:eventID/invited/teams", handler.GetInvitedTeam)
	e.GET("/event/:eventID/invitation/users", handler.GetInvitationUser)
	e.GET("/event/:eventID/invitation/teams", handler.GetInvitationTeam)
	return nil
}

func (eh *InviteHandler) Invite(ctx echo.Context) (err error) {
	inv := &models.Invitation{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, inv); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO хардкод
	inv.OwnerID, err = 1, nil
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	notify, err := eh.invite.Invite(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if notify {
		err = eh.notification.SendInviteNotification(*inv)
		if err != nil {
			log.Println("Notification wasnt sent: ", err)
		}
	}

	return nil
}

func (eh *InviteHandler) UnInvite(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	// TODO хардкод
	inv.OwnerID, err = 1, nil
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.invite.UnInvite(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (eh *InviteHandler) Deny(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	// TODO хардкод
	inv.OwnerID, err = 1, nil
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = eh.invite.Deny(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (eh *InviteHandler) IsInvited(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	// TODO хардкод
	inv.OwnerID, err = 1, nil
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.GuestID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	invited, err := eh.invite.IsInvited(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(models.IsInvited{IsInvited: invited}, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (eh *InviteHandler) GetInvitedUser(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	// TODO хардкод
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

	// TODO хардкод
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

	// TODO хардкод
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

	// TODO хардкод
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