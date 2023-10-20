package models

import (
	"github.com/lib/pq"
)

type UserPlants struct {
	UserID   uint64
	PlantsID pq.Int64Array `gorm:"type:integer[]"`
}

type Plant struct {
	UserID     uint64 `json:"userID"`
	ID         uint64 `json:"id"`
	ExternalID uint64 `json:"external_id"`
	CommonName string `json:"commonName"`
	ImageUrl   string `json:"imageUrl"`
}

type PerenualPlant struct {
	ID                       uint64    `json:"id"`
	Name                     string    `json:"common_name"`
	Family                   string    `json:"family"`
	Type                     string    `json:"type"`
	Cycle                    string    `json:"cycle"`
	Watering                 string    `json:"watering"`
	WateringPeriod           string    `json:"watering_period"`
	DepthWaterReq            UnitValue `json:"depth_water_requirement"`
	VolumeWaterReq           UnitValue `json:"volume_water_requirement"`
	WateringGeneralBenchmark UnitValue `json:"watering_general_benchmark"`
	PruningCount             struct {
		Amount   uint64 `json:"amount"`
		Interval string `json:"string"`
	} `json:"puring_count"`
	PruningMonth []string `json:"pruning_month"`
	Image        struct {
		URL string `json:"original_url"`
	} `json:"default_image"`
}

type UnitValue struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}
