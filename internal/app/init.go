package app

import (
	"strconv"

	httpAuth "github.com/cyber_bed/internal/auth/delivery"
	authMiddlewares "github.com/cyber_bed/internal/auth/delivery/middleware"
	authRepository "github.com/cyber_bed/internal/auth/repository"
	authUsecase "github.com/cyber_bed/internal/auth/usecase"
	"github.com/cyber_bed/internal/config"
	"github.com/cyber_bed/internal/domain"
	httpUsers "github.com/cyber_bed/internal/users/delivery"
	usersRepository "github.com/cyber_bed/internal/users/repository"
	usersUsecase "github.com/cyber_bed/internal/users/usecase"
	logger "github.com/cyber_bed/pkg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config

	usersUsecase domain.UsersUsecase
	authUsecase  domain.AuthUsecase

	usersHandler httpUsers.UsersHandler
	authHandler  httpAuth.AuthHandler

	authMiddleware *authMiddlewares.Middlewares
}

func New(e *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   e,
		Config: c,
	}
}

func (s *Server) init() {
	s.MakeUsecases()
	s.makeMiddlewares()
	s.MakeHandlers()
	s.MakeRouter()

	s.MakeEchoLogger()
}

func (s *Server) Start() error {
	s.init()
	return s.Echo.Start(
		s.Config.Server.Address + ":" + strconv.FormatUint(s.Config.Server.Port, 10),
	)
}

func (s *Server) MakeHandlers() {
	s.authHandler = httpAuth.NewAuthHandler(s.authUsecase, s.usersUsecase)
}

func (s *Server) MakeUsecases() {
	pgParams := s.Config.FormatDbAddr()

	authDB, err := authRepository.NewPostgres(pgParams)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	usersDB, err := usersRepository.NewPostgres(pgParams)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	s.authUsecase = authUsecase.NewAuthUsecase(authDB, usersDB)
	s.usersUsecase = usersUsecase.NewUsersUsecase(usersDB)
}

func (s *Server) MakeRouter() {
	v1 := s.Echo.Group("/api")
	v1.Use(logger.Middleware())
	v1.Use(middleware.Secure())

	v1.POST("/signup", s.authHandler.SignUp)
	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.authUsecase, s.usersUsecase)
}

func (s *Server) MakeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.SetLevel(logger.ToLevel(s.Config.LoggerLvl))
	s.Echo.Logger.Info("server started")
}
