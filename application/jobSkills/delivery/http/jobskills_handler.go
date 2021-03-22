package http

import (
	"diplomaProject/application/jobSkills"
	"diplomaProject/application/models"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type JobSkillsHandler struct {
	useCase jobskills.UseCase
}

func NewJobSkillsHandler(e *echo.Echo, usecase jobskills.UseCase) error {

	handler := JobSkillsHandler{useCase: usecase}

	e.GET("/job", handler.GetJobs)
	e.GET("/job/:name/skills", handler.GetSkillsByJob)
	e.POST("user/:id/skills", handler.AddSkill)
	e.DELETE("user/:id/skills", handler.RemoveSkill)

	return nil
}

func (js *JobSkillsHandler) GetJobs(ctx echo.Context) error {

	jArr, err := js.useCase.GetAllJobs()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(models.JobArr(*jArr), ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (js *JobSkillsHandler) GetSkillsByJob(ctx echo.Context) error {
	jobName, _ := url.QueryUnescape(ctx.Param("name"))

	sArr, err := js.useCase.GetSkillsByJob(jobName)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if _, err = easyjson.MarshalToWriter(models.SkillsArr(*sArr), ctx.Response().Writer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (js *JobSkillsHandler) AddSkill(ctx echo.Context) error {
	uID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	add := &models.AddSkillIDArr{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = js.useCase.AddSkill(uID, add)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return echo.NewHTTPError(200, "OK")
}

func (js *JobSkillsHandler) RemoveSkill(ctx echo.Context) error {
	uID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	add := &models.AddSkillID{}
	if err = easyjson.UnmarshalFromReader(ctx.Request().Body, add); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = js.useCase.RemoveSkill(uID, add.JobID, add.SkillID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	//if _, err = easyjson.MarshalToWriter(tm, ctx.Response().Writer); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	return echo.NewHTTPError(200, "OK")
}
