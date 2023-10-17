package domain

import (
	"context"
	"github.com/cyber_bed/internal/models"
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

type API interface {
	Recognize(
		ctx context.Context,
		formdata *multipart.Form,
		project models.Project,
	) ([]models.Plant, error)
}

type Usecase interface {
	Recognize(
		ctx context.Context,
		formdata *multipart.Form,
		project string,
	) ([]models.Plant, error)
}

type Handler interface {
	Recognize(c echo.Context) error
}
