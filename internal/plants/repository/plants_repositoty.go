package plantsRepository

import (
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cyber_bed/internal/models"
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

func (db *Postgres) CreateUserPlantsRelations(userID uint64, plantsID []int64) error {
	res := db.DB.Create(&models.UserPlants{
		UserID:   userID,
		PlantsID: pq.Int64Array(plantsID),
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *Postgres) AddUserPlantsRelations(userID uint64, plantsID []int64) error {
	userPlant := []models.UserPlants{}
	db.DB.Table(models.PlantsTable).Select("*").Where("user_id = ?", userID).Scan(&userPlant)

	if len(userPlant) == 0 {
		res := db.DB.Table(models.PlantsTable).Create(&models.UserPlants{
			UserID:   userID,
			PlantsID: pq.Int64Array(plantsID),
		})
		if res.Error != nil {
			return res.Error
		}
	} else {
		newPlantIDs := userPlant[0].PlantsID
		newPlantIDs = append(newPlantIDs, plantsID...)

		res := db.DB.Table(models.PlantsTable).Where("user_id = ?", userID).Update("plants_id", &newPlantIDs)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (db *Postgres) GetPlantsByID(userID uint64) (models.UserPlants, error) {
	var pl models.UserPlants
	if err := db.DB.Table(models.PlantsTable).
		Select("*").
		Where("user_id = ?", userID).
		Scan(&pl).
		Error; err != nil {
		return models.UserPlants{}, err
	}

	return pl, nil
}

func (db *Postgres) UpdateUserPlantsRelation(relation models.UserPlants) error {
	if err := db.DB.Table(models.PlantsTable).
		Where("user_id = ?", relation.UserID).
		Update("plants_id", &relation.PlantsID).Error; err != nil {
		return err
	}
	return nil
}
