package domain

type AuthRepository interface {
	CreateName(name string) error
}
