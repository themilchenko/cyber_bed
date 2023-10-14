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
	usr := models.Username{
		Username: user.Username,
	}

	res := db.DB.Table(models.UsersTable).Create(&usr)
	if res.Error != nil {
		return 0, res.Error
	}

	// Don't know what happend but if I throw usr.ID
	// in next sql query directly it's still not updated value
	// after first query and gorm throws error by this reason.
	// But when I check usr.ID while debuggind its correct.
	// So I assign urs.ID to user.ID to solve poblem.
	user.ID = usr.ID

	res = db.DB.Table(models.UsersInfoTable).Create(models.UsersInfo{
		UserID:   user.ID,
		Password: user.Password,
		Avatar:   user.Avatar,
	})
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}

func (db *Postgres) GetByUsername(username string) (models.User, error) {
	var usr models.User
	res := db.DB.Table(models.UsersTable).Where(&models.Username{
		Username: username,
	}).
		Select("*").
		Joins("JOIN users_info ON users_info.user_id=users.id").Scan(&usr)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return usr, nil
}

func (db *Postgres) GetByID(id uint64) (models.User, error) {
	var usr models.User
	res := db.DB.Table(models.UsersTable).Where(&models.User{
		ID: id,
	}).
		Select("*").
		Joins("JOIN users_info ON users_info.user_id=users.id").Scan(&usr)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return usr, nil
}

func (db *Postgres) GetBySessionID(sessionID string) (models.User, error) {
	var usr models.User
	res := db.DB.Table(models.SessionTable).
		Select("users.id, users.username, users_info.password, users_info.avatar").
		Where(&models.Cookie{Value: sessionID}).
		Joins("JOIN users_info ON sessions.user_id=users_info.user_id").
		Joins("JOIN users ON sessions.user_id=users.id").Scan(&usr)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	return usr, nil
}
