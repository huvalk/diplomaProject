package httpDebug

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"net/http"
)

type DebugHandler struct {
}

func NewDebugHandler(e *echo.Echo) error {
	handler := DebugHandler{}

	e.GET("/panic", handler.Panic)
	e.GET("/ping", handler.Ping)
	return nil
}

func (eh *DebugHandler) Panic(ctx echo.Context) error {
	panic("Test panic middleware")
}

func (eh *DebugHandler) Ping(ctx echo.Context) error {
	golog.Debug("Knok-knok")

	_, err := ctx.Response().Write([]byte("I am flying, Jack"))

	if err != nil {
		golog.Error("Knok failed: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
