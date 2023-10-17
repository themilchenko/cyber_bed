package authRepository

import (
	"github.com/cyber_bed/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
	}, nil
}

func (db *Postgres) CreateSession(cookie models.Cookie) (string, error) {
	res := db.DB.Table(models.SessionTable).Create(&cookie)
	if res.Error != nil {
		return "", res.Error
	}
	return cookie.Value, nil
}

func (db *Postgres) DeleteBySessionID(sessionID string) error {
	if err := db.DB.Table(models.SessionTable).
		Where("value = ?", sessionID).
		Delete(&models.Cookie{}).
		Error; err != nil {
		return err
	}
	return nil
}
