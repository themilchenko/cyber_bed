package usecase

import (
	"context"
	"github.com/cyber_bed/internal/models"
	domainPlantsAPI "github.com/cyber_bed/internal/plants-api"
	domainRecognition "github.com/cyber_bed/internal/recognize-api"
	"github.com/pkg/errors"
	"mime/multipart"
)

type usecase struct {
	apiRecognition domainRecognition.API
	apiPlants      domainPlantsAPI.PlantsAPI
}

func New(
	api domainRecognition.API,
	apiPlants domainPlantsAPI.PlantsAPI,
) domainRecognition.Usecase {
	return usecase{
		apiRecognition: api,
		apiPlants:      apiPlants,
	}
}

func (u usecase) Recognize(ctx context.Context, formdata *multipart.Form, project string) ([]models.Plant, error) {
	recognized, err := u.apiRecognition.Recognize(ctx, formdata, models.Project(project))
	if err != nil {
		return nil, errors.Wrap(err, "failed to recognize images")
	}

	plants := make([]models.Plant, 0)
	for _, plant := range recognized {
		found, err := u.apiPlants.Search(ctx, plant.CommonName)
		if err != nil {
			return nil, errors.Wrap(err, "failed to search plant")
		}

		plants = append(plants, found...)
	}

	return plants, nil
}
