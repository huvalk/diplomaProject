package http

import (
	"diplomaProject/application/jobSkills"
	"diplomaProject/application/models"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"net/url"
)

type JobSkillsHandler struct {
	useCase jobskills.UseCase
}

func NewJobSkillsHandler(e *echo.Echo, usecase jobskills.UseCase) error {

	handler := JobSkillsHandler{useCase: usecase}

	e.GET("/job", handler.GetJobs)
	e.GET("/job/:name/skills", handler.GetSkillsByJob)

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
