package models

import (
	"time"
)

type Cookie struct {
	UserID     uint64    `json:"value"`
	Value      string    `json:"userID"`
	ExpireDate time.Time `json:"expire_date"`
}
