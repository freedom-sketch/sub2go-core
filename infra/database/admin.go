package database

import (
	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func IsAdmin(db *gorm.DB, userUUID uuid.UUID) (bool, error) {
	var count int64
	err := db.Model(&models.Admin{}).
		Where("user_uuid = ?", userUUID.String()).Count(&count).Error
	return count > 0, err
}
