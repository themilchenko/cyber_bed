package plants_api

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/carlmjohnson/requests"
	"github.com/pkg/errors"

	"github.com/cyber_bed/internal/api/convert"
	"github.com/cyber_bed/internal/models"
)

type PlantsAPI interface {
	SearchByName(ctx context.Context, name string) ([]models.Plant, error)
	SearchByID(ctx context.Context, id uint64) (models.Plant, error)
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

func (t *TrefleAPI) SearchByName(
	ctx context.Context,
	name string,
) ([]models.Plant, error) {
	u := t.baseURL
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

	return convert.InputSearchResultsToModels(resp, t.countResults), nil
}

func (t *TrefleAPI) SearchByID(ctx context.Context, id uint64) (models.Plant, error) {
	u := t.baseURL
	q := u.Query()
	u.RawQuery = q.Encode()
	apiURL := u.JoinPath(strconv.FormatUint(id, 10)).String()

	var resp models.SearchResponse
	if err := requests.
		URL(apiURL).
		Method(http.MethodGet).
		ToJSON(&resp).
		Fetch(ctx); err != nil {
		return models.Plant{}, errors.Wrap(err, "failed to search plant by id")
	}
	return convert.SearchItemToPlantModel(resp.Data), nil
}
