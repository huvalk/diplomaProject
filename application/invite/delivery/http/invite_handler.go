package http

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/middleware"
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"errors"
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
		invite:       iu,
		notification: nu,
	}

	e.POST("/event/:eventID/user/:userID/invite", handler.Invite, middleware.UserID)
	e.POST("/event/:eventID/user/:userID/uninvite", handler.UnInvite, middleware.UserID)
	e.POST("/event/:eventID/user/:userID/decline", handler.Deny, middleware.UserID)
	e.POST("/event/:eventID/user/:userID/ban", handler.DenyAndBan, middleware.UserID)
	e.GET("/event/:eventID/invited/user/:userID", handler.IsInvited, middleware.UserID)
	e.GET("/event/:eventID/invited/users", handler.GetInvitedUser, middleware.UserID)
	e.GET("/event/:eventID/invited/teams", handler.GetInvitedTeam, middleware.UserID)
	e.GET("/event/:eventID/invitation/users", handler.GetInvitationUser, middleware.UserID)
	e.GET("/event/:eventID/invitation/teams", handler.GetInvitationTeam, middleware.UserID)
	return nil
}

func (eh *InviteHandler) Invite(ctx echo.Context) (err error) {
	inv := &models.Invitation{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, inv); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.OwnerID = userID

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

	inviters, invitees, err := eh.invite.Invite(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	err = eh.notification.SendInviteNotification(inviters, inv.EventID)
	if err != nil {
		log.Println("Notification wasnt sent: ", err)
	}
	err = eh.notification.SendInviteNotification(invitees, inv.EventID)
	if err != nil {
		log.Println("Notification wasnt sent: ", err)
	}

	return nil
}

func (eh *InviteHandler) UnInvite(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.OwnerID = userID

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

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.GuestID = userID

	inv.OwnerID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inviters, err := eh.invite.Deny(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	err = eh.notification.SendDenyNotification(inviters, inv.EventID)
	if err != nil {
		log.Println("Notification wasnt sent: ", err)
	}

	return nil
}

func (eh *InviteHandler) DenyAndBan(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.GuestID = userID

	inv.OwnerID, err = strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	inviters, err := eh.invite.DenyAndBan(inv)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	err = eh.notification.SendDenyNotification(inviters, inv.EventID)
	if err != nil {
		log.Println("Notification wasnt sent: ", err)
	}

	return nil
}

func (eh *InviteHandler) IsInvited(ctx echo.Context) (err error) {
	inv := &models.Invitation{}

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.OwnerID = userID

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

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.OwnerID = userID

	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	declined, err := strconv.ParseBool(ctx.QueryParam("declined"))
	if err != nil {
		err = nil
		declined = false
	}

	users, err := eh.invite.GetInvitedUser(inv, declined)
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

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.OwnerID = userID

	inv.EventID, err = strconv.Atoi(ctx.Param("eventID"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	declined, err := strconv.ParseBool(ctx.QueryParam("declined"))
	if err != nil {
		err = nil
		declined = false
	}

	teams, err := eh.invite.GetInvitedTeam(inv, declined)
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

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.GuestID = userID

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

	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	inv.GuestID = userID

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
