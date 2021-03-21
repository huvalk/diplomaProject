package httpDebug

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type DebugHandler struct {
}

func NewDebugHandler(e *echo.Echo) error {
	handler := DebugHandler{}

	e.GET("/redirect", handler.RedirectLogin)
	e.GET("/panic", handler.Panic)
	e.GET("/ping", handler.Ping)
	return nil
}

func (eh *DebugHandler) RedirectLogin(ctx echo.Context) error {
	scopeTemp := "account+email+photos"
	url := fmt.Sprintf("https://oauth.vk.com/authorize?response_type=code&" +
		"client_id=%s&redirect_uri=%s&scope=%s&state=%s", os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URI"),
		scopeTemp, os.Getenv("STATE"))

	return ctx.Redirect(http.StatusTemporaryRedirect, url)
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
