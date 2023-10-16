package domain

type AuthUsecase interface {
	SignUpByID(userID uint64) (string, error)
	Login(login, password string) (string, uint64, error)
	Auth(sessionID string) error
	Logout(sessionID string) error
}
