package middleware

import (
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Recovery struct{}

func NewRecoveryMiddleware() *Recovery {
	return &Recovery{}
}

var numbers = []rune("1234567890")

func genRequestNumber(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(s)
}

var rIDKey = "rID"

func formatPath(path string) string {
	pathArray := strings.Split(path[1:], "/")
	for i := range pathArray {
		if _, err := strconv.Atoi(pathArray[i]); err == nil {
			pathArray[i] = "*"
		}
	}
	return "/" + strings.Join(pathArray, "/")
}

func (m *Recovery) RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			rID, ok := r.Context().Value("rID").(string)

			err := recover()
			if err != nil {
				if ok {
					golog.Errorf("#%s Panic: %s", rID, err.(error).Error())
				} else {
					golog.Errorf("Panic with no id: %s", err.(error).Error())
				}

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal haha error",
				})

				c.Response().Writer.WriteHeader(http.StatusInternalServerError)
				c.Response().Writer.Write(jsonBody)
			}
		}()

		return next(c)
	}
}