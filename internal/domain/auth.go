package domain

type AuthUsecase interface {
	CreateName(name string) error
}
