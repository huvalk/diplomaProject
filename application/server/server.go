package server

import (
	httpAuth "diplomaProject/application/auth/delivery/http"
	repositoryAuth "diplomaProject/application/auth/repository"
	usecaseAuth "diplomaProject/application/auth/usecase"
	httpDebug "diplomaProject/application/debug/delivery/http"
	http2 "diplomaProject/application/event/delivery/http"
	repository2 "diplomaProject/application/event/repository"
	usecase2 "diplomaProject/application/event/usecase"
	http4 "diplomaProject/application/feed/delivery/http"
	repository4 "diplomaProject/application/feed/repository"
	usecase4 "diplomaProject/application/feed/usecase"
	httpInvite "diplomaProject/application/invite/delivery/http"
	repositoryInvite "diplomaProject/application/invite/repository"
	usecaseInvite "diplomaProject/application/invite/usecase"
	http5 "diplomaProject/application/jobSkills/delivery/http"
	repository5 "diplomaProject/application/jobSkills/repository"
	usecase5 "diplomaProject/application/jobSkills/usecase"
	httpNotification "diplomaProject/application/notification/delivery/http"
	repositoryNotification "diplomaProject/application/notification/repository"
	usecaseNotification "diplomaProject/application/notification/usecase"
	http3 "diplomaProject/application/team/delivery/http"
	repository3 "diplomaProject/application/team/repository"
	usecase3 "diplomaProject/application/team/usecase"
	"diplomaProject/application/user/delivery/http"
	"diplomaProject/application/user/repository"
	"diplomaProject/application/user/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

type Server struct {
	port string
	e    *echo.Echo
}

func NewServer(e *echo.Echo, db *pgxpool.Pool) *Server {
	//middleware WIP
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "https://team-up.online", "https://dev.team-up.online"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	//feed handler
	feeds := repository4.NewFeedDatabase(db)
	feed := usecase4.NewFeed(feeds)
	err := http4.NewFeedHandler(e, feed)
	if err != nil {
		log.Println(err)
		return nil
	}

	//notification
	notificationRepo := repositoryNotification.NewNotificationRepository(db)
	notificationUsecase := usecaseNotification.NewNotificationUsecase(notificationRepo)
	err = httpNotification.NewNotificationHandler(e, notificationUsecase)
	if err != nil {
		log.Println(err)
		return nil
	}

	//event handler
	events := repository2.NewEventDatabase(db)
	event := usecase2.NewEvent(events, feed, notificationUsecase)
	err = http2.NewEventHandler(e, event)
	if err != nil {
		log.Println(err)
		return nil
	}

	//invite repository
	invites := repositoryInvite.NewInviteRepository(db)

	//team handler
	teams := repository3.NewTeamDatabase(db)
	team := usecase3.NewTeam(teams, events, notificationUsecase, invites)
	err = http3.NewTeamHandler(e, team)
	if err != nil {
		log.Println(err)
		return nil
	}

	//user handler
	//sessions := session.NewSessionDatabase(rd)
	users := repository.NewUserDatabase(db)
	user := usecase.NewUser(users, feeds, team)
	err = http.NewUserHandler(e, user)
	if err != nil {
		log.Println(err)
		return nil
	}

	//jobskills handler
	jobs := repository5.NewJobSkillsDatabase(db)
	job := usecase5.NewJobSkills(jobs)
	err = http5.NewJobSkillsHandler(e, job)
	if err != nil {
		log.Println(err)
		return nil
	}

	//invite handler
	invite := usecaseInvite.NewInviteUseCase(invites, users, teams)
	err = httpInvite.NewInviteHandler(e, invite, notificationUsecase)
	if err != nil {
		log.Println(err)
		return nil
	}

	//debug
	err = httpDebug.NewDebugHandler(e)
	if err != nil {
		log.Println(err)
		return nil
	}

	//auth
	authRepo := repositoryAuth.NewAuthRepository(db)
	authUsecase := usecaseAuth.NewUsecase(authRepo, users, events)
	err = httpAuth.NewAuthHandler(e, authUsecase)
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
