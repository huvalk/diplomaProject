package server

import (
	http2 "diplomaProject/application/event/delivery/http"
	repository2 "diplomaProject/application/event/repository"
	usecase2 "diplomaProject/application/event/usecase"
	http4 "diplomaProject/application/feed/delivery/http"
	repository4 "diplomaProject/application/feed/repository"
	usecase4 "diplomaProject/application/feed/usecase"
	http3 "diplomaProject/application/team/delivery/http"
	repository3 "diplomaProject/application/team/repository"
	usecase3 "diplomaProject/application/team/usecase"
	"diplomaProject/application/user/delivery/http"
	"diplomaProject/application/user/repository"
	"diplomaProject/application/user/usecase"
	"github.com/labstack/echo"
	"log"
)

type Server struct {
	port string
	e    *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	//middleware WIP

	//feed handler
	feeds := repository4.NewFeedDatabase(nil)
	feed := usecase4.NewFeed(feeds)
	err := http4.NewFeedHandler(e, feed)
	if err != nil {
		log.Println(err)
		return nil
	}

	//user handler
	//sessions := session.NewSessionDatabase(rd)
	users := repository.NewUserDatabase(nil)
	user := usecase.NewUser(users, feeds)
	err = http.NewUserHandler(e, user)
	if err != nil {
		log.Println(err)
		return nil
	}

	//event handler
	events := repository2.NewEventDatabase(nil)
	event := usecase2.NewEvent(events, feeds)
	err = http2.NewEventHandler(e, event)
	if err != nil {
		log.Println(err)
		return nil
	}

	//team handler
	teams := repository3.NewTeamDatabase(nil)
	team := usecase3.NewTeam(teams, events)
	err = http3.NewTeamHandler(e, team)
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

func (s Server) ListenAndServeTLS(sslPath string) error {
	return s.e.StartTLS(s.port, sslPath+"/fullchain.pem", sslPath+"/privkey.pem")
}
