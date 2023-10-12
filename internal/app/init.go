package app

import (
	"cyber_bed/internal/config"
	logger "cyber_bed/pkg"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

func New(e *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   e,
		Config: c,
	}
}

func (s *Server) init() {
	s.MakeEchoLogger()
}

func (s *Server) Start() error {
	s.init()
	return s.Echo.Start(
		s.Config.Server.Address + ":" + strconv.FormatUint(s.Config.Server.Port, 10),
	)
}

func (s *Server) MakeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.SetLevel(logger.ToLevel(s.Config.LoggerLvl))
	s.Echo.Logger.Info("server started")
}
