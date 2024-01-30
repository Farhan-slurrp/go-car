package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := `postgres://postgres.vfofoxaunwukvggzbwoq:S4p3r$dmin~@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
