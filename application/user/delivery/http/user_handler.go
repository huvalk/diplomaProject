package http

import (
	"diplomaProject/application/user"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"strconv"
)

type UserHandler struct {
	useCase user.UseCase
}

func NewUserHandler(e *echo.Echo, usecase user.UseCase) error {

	handler := UserHandler{useCase: usecase}

	e.GET("/oauth", handler.Hello)
	e.GET("/user/:id", handler.Profile)
	return nil
}

func (uh UserHandler) Hello(ctx echo.Context) error {
	return ctx.String(200, "WORKS!")
}

func (uh *UserHandler) Profile(ctx echo.Context) error {
	uid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return err
	}
	usr, err := uh.useCase.Get(uid)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.String(200, fmt.Sprintf("%v", *usr))
}
