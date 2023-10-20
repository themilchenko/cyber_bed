package plants_api

import (
	"context"
	"net/url"

	"github.com/cyber_bed/internal/domain"
	"github.com/cyber_bed/internal/models"
)

type PerenualAPI struct {
	baseURL *url.URL
}

func NewPerenualAPI(url *url.URL, token string) domain.PlantsAPI {
	url.Query().Set("key", token)
	q := url.Query()
	q.Set("key", token)

	url.RawQuery = q.Encode()
	return &PerenualAPI{
		baseURL: url,
	}
}

func (p *PerenualAPI) SearchByName(ctx context.Context, name string) ([]models.Plant, error) {
	return nil, nil
}

func (p *PerenualAPI) SearchByID(ctx context.Context, id uint64) (models.Plant, error) {
	return models.Plant{}, nil
}

func (p *PerenualAPI) GetPage(ctx context.Context, pageNum uint64) ([]models.Plant, error) {
	return nil, nil
}
