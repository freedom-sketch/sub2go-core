package models

import "time"

type Server struct {
	ID          uint      `gorm:"primarykey;autoIncrement"`
	Name        string    `gorm:"uniqueIndex;not null"`
	Host        string    `gorm:"size:255;not null"`
	APIAddress  string    `gorm:"size:255;not null"`
	APIKey      string    `gorm:"size:255;not null;uniqueIndex"`
	Status      string    `gorm:"size:255;default:'offline'"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true;index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Inbounds []Inbound `gorm:"foreignKey:ServerID;constraint:OnDelete:CASCADE"`
}

func (Server) TableName() string {
	return "servers"
}
