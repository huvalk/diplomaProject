package http

import (
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"io/ioutil"
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
	e.GET("/event/:evtid/user/:id/team", handler.GetTeamByUser)
	e.POST("/event/:evtid/team", handler.CreateTeam)
	e.POST("/team/:id/add", handler.AddMember)
	e.POST("/event/:evtid/team/join", handler.Union)
	return nil
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

	var body []byte
	defer ctx.Request().Body.Close()
	body, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newTeam := &models.Team{}
	err = newTeam.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
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

func (th *TeamHandler) AddMember(ctx echo.Context) error {
	tID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var body []byte
	defer ctx.Request().Body.Close()
	body, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	add := &models.AddToTeam{}
	err = add.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
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

	var body []byte
	defer ctx.Request().Body.Close()
	body, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	add := &models.AddToUser{}
	err = add.UnmarshalJSON(body)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusConflict, err.Error())
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
