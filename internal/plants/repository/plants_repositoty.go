package plantsRepository

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

func (db *Postgres) CreateUserPlantsRelations(userID uint64, plantsID []uint64) error {
	res := db.DB.Create(&models.UserPlants{
		UserID:   userID,
		PlantIDs: plantsID,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *Postgres) AddUserPlantsRelations(userID uint64, plantsID []uint64) error {
	var userPlant models.UserPlants
	if db.DB.Where("user_id = ?", userID).First(&userPlant).Error != nil {
		userPlant = models.UserPlants{
			UserID:   userID,
			PlantIDs: plantsID,
		}
		res := db.DB.Create(&userPlant)
		if res.Error != nil {
			return res.Error
		}
	} else {
		newPlantIDs := append(userPlant.PlantIDs, plantsID...)
		res := db.DB.Model(&userPlant).Update("plants_id", newPlantIDs)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (db *Postgres) GetPlantsByID(userID uint64) ([]uint64, error) {
	var plantsID []uint64
	if err := db.DB.Table(models.PlantsTable).
		Where("user_id = ?", userID).
		Pluck("plants_id", &plantsID).
		Error; err != nil {
		return nil, err
	}
	return plantsID, nil
}
