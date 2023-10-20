package httpPlants

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	httpAuth "github.com/cyber_bed/internal/auth/delivery"
	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
)

type PlantsHandler struct {
	plantsUsecase domain.PlantsUsecase
	usersUsecase  domain.UsersUsecase
	trefleAPI     domain.PlantsAPI
}

func NewPlantsHandler(
	p domain.PlantsUsecase,
	u domain.UsersUsecase,
	pl domain.PlantsAPI,
) PlantsHandler {
	return PlantsHandler{
		plantsUsecase: p,
		usersUsecase:  u,
		trefleAPI:     pl,
	}
}

// ================================================
// Handlers for handling requests with external API
// ================================================

func (h PlantsHandler) GetPlantFromAPI(c echo.Context) error {
	plantID, err := strconv.ParseUint(c.Param("plantID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	plant, err := h.trefleAPI.SearchByID(c.Request().Context(), plantID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, plant)
}

func (h PlantsHandler) GetPlantsFromAPI(c echo.Context) error {
	pageNum, err := strconv.ParseUint(c.QueryParam("page"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	plants, err := h.trefleAPI.GetPage(c.Request().Context(), pageNum)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, plants)
}

// ===============================================
// Handlers for handling authorized users requests
// ===============================================

func (h PlantsHandler) CreatePlant(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.usersUsecase.GetUserIDBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	plantID, err := strconv.ParseUint(c.Param("plantID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// TODO:
	// Need to bind additional information about plant
	// For example: date of last watering
	var recievedPlant models.Plant
	if err := c.Bind(&recievedPlant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	recievedPlant.ID = plantID
	recievedPlant.UserID = userID

	if err := h.plantsUsecase.AddPlant(recievedPlant); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, models.EmptyModel{})
}

func (h PlantsHandler) GetPlant(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.usersUsecase.GetUserIDBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	plantID, err := strconv.ParseUint(c.Param("plantID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	plant, err := h.plantsUsecase.GetPlant(userID, int64(plantID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	plant, err = h.trefleAPI.SearchByID(c.Request().Context(), plant.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, plant)
}

func (h PlantsHandler) GetPlants(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.usersUsecase.GetUserIDBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	plants, err := h.plantsUsecase.GetPlants(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for index, pl := range plants {
		plants[index], err = h.trefleAPI.SearchByID(c.Request().Context(), pl.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
	}

	return c.JSON(http.StatusOK, plants)
}

func (h PlantsHandler) DeletePlant(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	userID, err := h.usersUsecase.GetUserIDBySessionID(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	plantID, err := strconv.ParseUint(c.Param("plantID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err = h.plantsUsecase.DeletePlant(userID, plantID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, models.EmptyModel{})
}
