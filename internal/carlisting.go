package internal

import "gorm.io/gorm"

type CarListing struct {
	gorm.Model
	CarModel      string  `json:"model"`
	DailyPrice    float64 `json:"dailyPrice"`
	AvailableFrom string  `json:"availableFrom,omitempty"`
	AvailableTo   string  `json:"availableTo,omitempty"`
	OwnerId       string  `json:"ownerId"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
