package models

import (
	"time"
)

type Cookie struct {
	UserID     uint64    `json:"value"       gorm:"foreignkey:User"`
	Value      string    `json:"userID"`
	ExpireDate time.Time `json:"expire_date"`
}
