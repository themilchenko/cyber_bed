package usecase

import (
	"context"
	"mime/multipart"

	"github.com/pkg/errors"

	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
	domainRecognition "github.com/cyber_bed/internal/recognize-api"
)

type usecase struct {
	apiRecognition domainRecognition.API
	apiPlants      domain.PlantsAPI
}

func New(
	api domainRecognition.API,
	apiPlants domain.PlantsAPI,
) domainRecognition.Usecase {
	return usecase{
		apiRecognition: api,
		apiPlants:      apiPlants,
	}
}

func (u usecase) Recognize(
	ctx context.Context,
	formdata *multipart.Form,
	project string,
) ([]models.Plant, error) {
	recognized, err := u.apiRecognition.Recognize(ctx, formdata, models.Project(project))
	if err != nil {
		return nil, errors.Wrap(err, "failed to recognize images")
	}

	plants := make([]models.Plant, 0)
	for _, plant := range recognized {
		found, err := u.apiPlants.SearchByName(ctx, plant.CommonName)
		if err != nil {
			return nil, errors.Wrap(err, "failed to search plant")
		}

		plants = append(plants, found...)
	}

	return plants, nil
}
