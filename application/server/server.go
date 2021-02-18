package server

import (
	"diplomaProject/application/user/delivery/http"
	"diplomaProject/application/user/repository"
	user2 "diplomaProject/application/user/usecase"
	"github.com/labstack/echo"
	"log"
)

type Server struct {
	port string
	e    *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	//middleware WIP

	//user handler
	//sessions := session.NewSessionDatabase(rd)
	users := repository.NewUserDatabase(nil)
	user := user2.NewUser(users)
	err := http.NewUserHandler(e, user)
	if err != nil {
		log.Println(err)
		return nil
	}

	//prometeus

	//prometheus.MustRegister(middleware.FooCount, middleware.Hits)
	//e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return &Server{
		port: ":8080",
		e:    e,
	}
}

func (s Server) ListenAndServe() error {
	return s.e.Start(s.port)
}
