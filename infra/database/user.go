package database

import (
	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
