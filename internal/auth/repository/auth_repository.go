package authRepository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
	}, nil
}

func (db *Postgres) Close() error {
	return db.DB.Close()
}

func (db *Postgres) CreateName(name string) error {
	_, err := db.DB.Exec("INSERT INTO names (name) VALUES ($1)", name)
	if err != nil {
		return err
	}
	return nil
}
