package models

import "time"

type Subscription struct {
	ID           uint      `gorm:"primarykey;autoIncrement"`
	UserUUID     string    `gorm:"type:uuid;index;constraint:OnDelete:CASCADE"`
	Email        string    `gorm:"size:128;uniqueIndex"`
	TotalTraffic int64     `gorm:"default:0"`
	UsedTraffic  int64     `gorm:"default:0"`
	StartDate    time.Time `gorm:"not null;index"`
	EndDate      time.Time `gorm:"not null;index"`
	IsActive     bool      `gorm:"default:true;index"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserUUID"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
