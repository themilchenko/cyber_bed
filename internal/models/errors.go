package models

import "github.com/pkg/errors"

var (
	ErrNoImages      = errors.New("no images provided")
	ErrTooManyImages = errors.New("too many images")
	ErrUserExists    = errors.New("user already exists")
)
