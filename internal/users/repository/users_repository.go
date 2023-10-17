package usersRepository

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

func (db *Postgres) Create(user models.User) (uint64, error) {
	var usr models.Username

	err := db.DB.Table(models.UsersTable).Create(&models.Username{
		Username: user.Username,
	}).Scan(&usr).Error
	if err != nil {
		return 0, err
	}

	err = db.DB.Table(models.UsersInfoTable).Create(models.UsersInfo{
		UserID:   usr.ID,
		Password: user.Password,
		Avatar:   user.Avatar,
	}).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (db *Postgres) GetByUsername(username string) (models.User, error) {
	var usr models.User
	err := db.DB.Table(models.UsersTable).
		Select("users.id, users.username, users_info.password, users_info.avatar").
		Joins("JOIN users_info ON users.id=users_info.user_id").
		Where("users.username = ?", username).Last(&usr).Error
	if err != nil {
		return models.User{}, err
	}
	return usr, nil
}

func (db *Postgres) GetByID(id uint64) (models.User, error) {
	var usr models.User
	err := db.DB.Table(models.UsersTable).Where(&models.User{
		ID: id,
	}).
		Select("*").
		Joins("JOIN users_info ON users_info.user_id=users.id").Scan(&usr).Error
	if err != nil {
		return models.User{}, err
	}
	return usr, nil
}

func (db *Postgres) GetUserIDBySessionID(sessionID string) (uint64, error) {
	var usrID uint64
	if err := db.DB.Table(models.SessionTable).
		Select("user_id").
		Where("value = ?", sessionID).
		Scan(&usrID).Error; err != nil {
		return 0, err
	}
	return usrID, nil
}

func (db *Postgres) GetBySessionID(sessionID string) (models.User, error) {
	var usr models.User
	err := db.DB.Table(models.SessionTable).
		Select("users.id, users.username, users_info.password, users_info.avatar").
		Where(&models.Cookie{Value: sessionID}).
		Joins("JOIN users_info ON sessions.user_id=users_info.user_id").
		Joins("JOIN users ON sessions.user_id=users.id").Scan(&usr).Error
	if err != nil {
		return models.User{}, err
	}
	return usr, nil
}
