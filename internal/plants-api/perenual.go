package plants_api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/carlmjohnson/requests"
	"github.com/pkg/errors"

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
	u := p.baseURL
	q := u.Query()
	q.Set("q", name)

	u.RawQuery = q.Encode()
	apiURL := u.JoinPath("search")

	var resp models.SearchSliceResponse

	if err := requests.
		URL(apiURL.String()).
		Method(http.MethodGet).
		ToJSON(&resp).
		Fetch(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to search plant by name")
	}

	return nil, nil
}

func (p *PerenualAPI) SearchByID(ctx context.Context, id uint64) (models.Plant, error) {
	return models.Plant{}, nil
}

func (p *PerenualAPI) GetPage(ctx context.Context, pageNum uint64) ([]models.Plant, error) {
	return nil, nil
}
