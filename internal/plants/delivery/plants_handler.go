package httpPlants

import (
	"net/http"
	"strconv"

	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
	"github.com/labstack/echo/v4"
)

type PlantsHandler struct {
	plantsUsecase domain.PlantsUsecase
}

func NewPlantsHandler(p domain.PlantsUsecase) PlantsHandler {
	return PlantsHandler{
		plantsUsecase: p,
	}
}

func (h PlantsHandler) CreatePlant(c echo.Context) error {
	var recievedPlant models.Plant
	if err := c.Bind(&recievedPlant); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.plantsUsecase.AddPlant(recievedPlant); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, []interface{}{})
}

func (h PlantsHandler) GetPlant(c echo.Context) error {
	return nil
}

func (h PlantsHandler) GetPlants(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	plants, err := h.plantsUsecase.GetPlants(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, plants)
}
