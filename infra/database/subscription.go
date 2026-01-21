package database

import (
	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HasActiveSubscription checks if the user with the given UUID has an active subscription
func HasActiveSubscription(db *gorm.DB, userUUID uuid.UUID) (bool, error) {
	var count int64
	err := db.Model(&models.Subscription{}).
		Where("user_uuid = ? AND is_active = true AND end_date > NOW()", userUUID.String()).
		Count(&count).Error
	return count > 0, err
}

// GetSubscriptionByUserUUID retrieves the subscription record for the given user UUID
func GetSubscriptionByUserUUID(db *gorm.DB, userUUID uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	err := db.Where("user_uuid = ?", userUUID.String()).First(&subscription).Error
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}
