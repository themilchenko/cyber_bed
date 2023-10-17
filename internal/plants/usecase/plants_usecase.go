package plantsUsecase

import (
	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
)

type PlantsUsecase struct {
	plantsRepository domain.PlantsRepository
}

func NewPlansUsecase(p domain.PlantsRepository) domain.PlantsUsecase {
	return PlantsUsecase{
		plantsRepository: p,
	}
}

func (u PlantsUsecase) AddPlant(plant models.Plant) error {
	if err := u.plantsRepository.AddUserPlantsRelations(plant.UserID, []int64{int64(plant.ID)}); err != nil {
		return err
	}
	return nil
}

func (u PlantsUsecase) GetPlant(userID uint64, plantID int64) (models.Plant, error) {
	return models.Plant{}, nil
}

func (u PlantsUsecase) GetPlants(userID uint64) ([]models.Plant, error) {
	plantsIDs, err := u.plantsRepository.GetPlantsByID(userID)
	if err != nil {
		return nil, err
	}

	pl := plantsIDs.PlantsID
	plants := make([]models.Plant, 0)
	for _, p := range pl {
		plants = append(plants, models.Plant{
			ID: uint64(p),
		})
	}

	// Here we will go to plants service

	return plants, nil
}
