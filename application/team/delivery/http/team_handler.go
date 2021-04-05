package http

import (
	"diplomaProject/application/middleware"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"errors"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	useCase team.UseCase
}

func NewTeamHandler(e *echo.Echo, usecase team.UseCase) error {

	handler := TeamHandler{useCase: usecase}

	e.GET("/team/:id", handler.GetTeam)
	e.POST("/team/:id", handler.UpdateTeam)
	e.POST("/team/:id/vote", handler.SendVote, middleware.UserID)
	e.GET("/event/:evtid/user/:id/team", handler.GetTeamByUser)
	e.POST("/event/:evtid/team", handler.CreateTeam)
	e.POST("/team/:id/add", handler.AddMember)
	e.POST("/team/:id/leave", handler.Leave, middleware.UserID)
	e.POST("/event/:evtid/team/join", handler.Union, middleware.UserID)
	return nil
}

func (th *TeamHandler) SendVote(ctx echo.Context) error {
	tID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	vt := &models.Vote{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, vt); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	vt.TeamID = tID
	vt.WhoID = userID
	//if userID != add.UID {
	//	return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	//}

	err = th.useCase.SendVote(vt)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return ctx.String(200, "OK")
}

func (th *TeamHandler) GetTeamByUser(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tm, err := th.useCase.GetTeamByUser(uid, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) GetTeam(ctx echo.Context) error {
	tid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tm, err := th.useCase.Get(tid)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) CreateTeam(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTeam := &models.Team{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, newTeam); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTeam, err = th.useCase.Create(newTeam, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(newTeam, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) UpdateTeam(ctx echo.Context) error {
	tmID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTeam := &models.Team{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, newTeam); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newTeam.Id = tmID
	newTeam, err = th.useCase.SetName(newTeam)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(newTeam, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) Leave(ctx echo.Context) error {
	tID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	add := &models.AddToTeam{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	if userID != add.UID {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	}

	tm, err := th.useCase.RemoveMember(tID, add.UID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) AddMember(ctx echo.Context) error {
	tID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	add := &models.AddToTeam{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// TODO не знаю, как проверить оригинального пользователя

	tm, err := th.useCase.AddMember(tID, add.UID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (th *TeamHandler) Union(ctx echo.Context) error {
	evtID, err := strconv.Atoi(ctx.Param("evtid"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	add := &models.AddToUser{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, found := ctx.Get("userID").(int)
	if !found {
		log.Println("userID not found")
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("userID not found"))
	}
	if userID != add.UID1 && userID != add.UID2 {
		return echo.NewHTTPError(http.StatusUnauthorized, errors.New("userID doesnt match current user"))
	}

	tm, err := th.useCase.Union(add.UID1, add.UID2, evtID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
