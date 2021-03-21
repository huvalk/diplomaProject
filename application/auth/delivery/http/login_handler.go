package httpAuth

import (
	"diplomaProject/application/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
}

func NewAuthHandler(e *echo.Echo) error {
	handler := AuthHandler{}

	//e.POST("/login", handler.Login)
	e.GET("/auth", handler.Auth)
	return nil
}

func (eh *AuthHandler) Auth(ctx echo.Context) error {
	stateTemp := ctx.QueryParam("state")
	if stateTemp == "" || stateTemp != os.Getenv("STATE") {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("state doesnt match"))
	}
	code := ctx.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("no code provided"))
	}
	url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&" +
		"redirect_uri=%s&client_id=%s&client_secret=%s", code, os.Getenv("REDIRECT_URI"),
		os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()
	token := struct {
		AccessToken string `json:"access_token"`
	}{}
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &token)
	url = fmt.Sprintf("https://api.vk.com/method/%s?v=5.124&access_token=%s", "users.get", token.AccessToken)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	golog.Error(string(bytes))

	return nil
}

func (eh *AuthHandler) Login(ctx echo.Context) error {
	user := &models.AuthUser{}
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, user); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Time{},
		MaxAge:   1000000,
		Secure:   false,
		HttpOnly: false,
	})
	return nil
}
