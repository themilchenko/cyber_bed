package plantsUsecase

import "github.com/cyber_bed/internal/domain"

type PlantsUsecase struct {
	plantsRepository domain.PlantsRepository
}

func NewPlansUsecase(p domain.PlantsRepository) domain.PlantsUsecase {
	return PlantsUsecase{
		plantsRepository: p,
	}
}
