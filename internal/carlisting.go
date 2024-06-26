package internal

import (
	"time"

	"gorm.io/gorm"
)

type CarListing struct {
	gorm.Model
	CarModel      string    `json:"car_model"`
	DailyPrice    float64   `json:"daily_price"`
	AvailableFrom time.Time `json:"available_from,omitempty"`
	AvailableTo   time.Time `json:"available_to,omitempty"`
	OwnerId       uint      `json:"owner_id"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
