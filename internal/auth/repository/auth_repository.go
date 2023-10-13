package authRepository

import (
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

// Only for test
func (db *Postgres) CreateName(name string) error {
	type Name struct {
		Name string
	}
	res := db.DB.Table("names").Create(&Name{
		Name: name,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
