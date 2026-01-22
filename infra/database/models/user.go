package models

import "time"

type User struct {
	UUID       string    `gorm:"type:uuid;primarykey"`
	TelegramID *int64    `gorm:"column:tg_id;uniqueIndex"`
	IsActive   bool      `gorm:"default:true;index"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Subscriptions []Subscription `gorm:"foreignKey:UserUUID;constraint:OnDelete:CASCADE"`
	Admin         *Admin         `gorm:"foreignKey:UserUUID;references:UUID"`
}

func (User) TableName() string {
	return "users"
}
