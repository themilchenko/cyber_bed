package app

import (
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	httpAuth "github.com/cyber_bed/internal/auth/delivery"
	authMiddlewares "github.com/cyber_bed/internal/auth/delivery/middleware"
	authRepository "github.com/cyber_bed/internal/auth/repository"
	authUsecase "github.com/cyber_bed/internal/auth/usecase"
	"github.com/cyber_bed/internal/config"
	"github.com/cyber_bed/internal/domain"
	domainPlantsAPI "github.com/cyber_bed/internal/plants-api"
	httpPlants "github.com/cyber_bed/internal/plants/delivery"
	plantsRepository "github.com/cyber_bed/internal/plants/repository"
	plantsUsecase "github.com/cyber_bed/internal/plants/usecase"
	domainRecognition "github.com/cyber_bed/internal/recognize-api"
	"github.com/cyber_bed/internal/recognize-api/delivery/http"
	recAPI "github.com/cyber_bed/internal/recognize-api/repository/api"
	recUsecase "github.com/cyber_bed/internal/recognize-api/usecase"
	httpUsers "github.com/cyber_bed/internal/users/delivery"
	usersRepository "github.com/cyber_bed/internal/users/repository"
	usersUsecase "github.com/cyber_bed/internal/users/usecase"
	logger "github.com/cyber_bed/pkg"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config

	usersUsecase  domain.UsersUsecase
	authUsecase   domain.AuthUsecase
	plantsUsecase domain.PlantsUsecase
	recUsecase    domainRecognition.Usecase
	plantsAPI     domainPlantsAPI.PlantsAPI

	usersHandler  httpUsers.UsersHandler
	authHandler   httpAuth.AuthHandler
	recHandler    domainRecognition.Handler
	plantsHandler httpPlants.PlantsHandler

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
	s.recHandler = http.NewHandler(s.recUsecase)
	s.authHandler = httpAuth.NewAuthHandler(s.authUsecase, s.usersUsecase, s.Config.CookieSettings)
	s.plantsHandler = httpPlants.NewPlantsHandler(s.plantsUsecase, s.usersUsecase)
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

	plantsDB, err := plantsRepository.NewPostgres(pgParams)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	u, err := url.Parse(s.Config.RecognizeAPI.BaseURL)
	if err != nil {
		s.Echo.Logger.Error(errors.Wrap(err, "failed to parse base url"))
		return
	}

	recognitionAPI := recAPI.NewRecognitionAPI(
		u,
		s.Config.RecognizeAPI.Token,
		s.Config.RecognizeAPI.ImageField,
		s.Config.RecognizeAPI.MaxImages,
		s.Config.RecognizeAPI.CountResults,
	)

	if u, err = url.Parse(s.Config.TrefleAPI.BaseURL); err != nil {
		s.Echo.Logger.Error(errors.Wrap(err, "failed to parse base url"))
		return
	}

	s.plantsAPI = domainPlantsAPI.NewTrefleAPI(
		u,
		s.Config.TrefleAPI.CountPlants,
		s.Config.TrefleAPI.Token,
	)

	s.authUsecase = authUsecase.NewAuthUsecase(authDB, usersDB, s.Config.CookieSettings)
	s.usersUsecase = usersUsecase.NewUsersUsecase(usersDB)
	s.plantsUsecase = plantsUsecase.NewPlansUsecase(plantsDB, s.plantsAPI)
	s.recUsecase = recUsecase.New(recognitionAPI, s.plantsAPI)
}

func (s *Server) MakeRouter() {
	v1 := s.Echo.Group("/api")
	v1.Use(logger.Middleware())
	v1.Use(middleware.Secure())

	v1.POST("/signup", s.authHandler.SignUp)
	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)

	v1.POST("/recognize", s.recHandler.Recognize)

	plants := v1.Group("/plants", s.authMiddleware.LoginRequired)
	plants.GET("/:plantID", s.plantsHandler.GetPlant)
	plants.POST("/:plantID", s.plantsHandler.CreatePlant)
	plants.DELETE("/:plantID", s.plantsHandler.DeletePlant)
	plants.GET("", s.plantsHandler.GetPlants)
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.authUsecase, s.usersUsecase)
}

func (s *Server) MakeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.SetLevel(logger.ToLevel(s.Config.LoggerLvl))
	s.Echo.Logger.Info("server started")
}
