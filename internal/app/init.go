package app

import (
	httpAuth "cyber_bed/internal/auth/delivery"
	authRepository "cyber_bed/internal/auth/repository"
	authUsecase "cyber_bed/internal/auth/usecase"
	"cyber_bed/internal/config"
	"cyber_bed/internal/domain"
	logger "cyber_bed/pkg"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config

	authUsecase domain.AuthUsecase

	authHandler httpAuth.AuthHandler
}

func New(e *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   e,
		Config: c,
	}
}

func (s *Server) init() {
	s.MakeUsecases()
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
	s.authHandler = httpAuth.NewAuthHandler(s.authUsecase)
}

func (s *Server) MakeUsecases() {
	pgParams := s.Config.FormatDbAddr()
	authDB, err := authRepository.NewPostgres(pgParams)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	s.authUsecase = authUsecase.NewAuthUsecase(authDB)
}

func (s *Server) MakeRouter() {
	v1 := s.Echo.Group("/api")
	v1.Use(logger.Middleware())
	v1.Use(middleware.Secure())

	v1.GET("/hello/:name", s.authHandler.CreateName)
}

func (s *Server) MakeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.SetLevel(logger.ToLevel(s.Config.LoggerLvl))
	s.Echo.Logger.Info("server started")
}
