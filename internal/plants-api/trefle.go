package plants_api

import (
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/cyber_bed/internal/api/convert"
	"github.com/cyber_bed/internal/models"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type PlantsAPI interface {
	Search(ctx context.Context, name string) ([]models.Plant, error)
}

type TrefleAPI struct {
	baseURL      *url.URL
	countResults int
}

func NewTrefleAPI(baseURL *url.URL, countResults int, token string) PlantsAPI {
	baseURL.Query().Set("token", token)
	q := baseURL.Query()
	q.Set("token", token)

	baseURL.RawQuery = q.Encode()

	return &TrefleAPI{
		baseURL:      baseURL,
		countResults: countResults,
	}
}

func (t *TrefleAPI) Search(
	ctx context.Context,
	name string,
) ([]models.Plant, error) {
	u := t.baseURL
	q := u.Query()
	q.Set("q", name)

	u.RawQuery = q.Encode()
	apiURL := u.JoinPath("search")

	var resp models.SearchResponse

	if err := requests.
		URL(apiURL.String()).
		Method(http.MethodGet).
		ToJSON(&resp).
		Fetch(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to search plant by name")
	}

	return convert.InputSearchResultsToModels(resp, t.countResults), nil
}
