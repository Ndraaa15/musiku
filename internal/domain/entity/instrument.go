package entity

import (
	"time"
)

type Instrument struct {
	ID            uint      `json:"id" gorm:"autoIncreament;primaryKey"`
	Name          string    `json:"name"`
	RentPrice     float64   `json:"rent_price"`
	Address       string    `json:"address"`
	Description   string    `json:"description"`
	Spesification string    `json:"spesification"`
	Status        bool      `json:"status"`
	CreateAt      time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time `json:"update_at" gorm:"autoUpdateTime"`
}
