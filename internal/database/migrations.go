package database

import (
	"github.com/sub2freedom/xray-manager/internal/database/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Subscription{},
		&models.Server{},
		&models.Inbound{},
	)
}
