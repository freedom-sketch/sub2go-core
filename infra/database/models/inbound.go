package models

import "time"

type Inbound struct {
	ID          uint   `gorm:"primarykey;autoIncrement"`
	ServerID    uint   `gorm:"index;constraint:OnDelete:CASCADE"`
	Tag         string `gorm:"size:255;uniqueIndex;not null"`
	Protocol    string `gorm:"size:64;not null"`
	Port        int    `gorm:"not null;index"`
	Network     string `gorm:"size:64;not null"`
	Security    string `gorm:"size:64;default:'reality'"`
	Flow        string `gorm:"size:64;default:''"`
	ShortIds    string `gorm:"type:text"`
	PublicKey   string
	PrivateKey  string
	Target      string
	SNI         string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true;index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Server Server `gorm:"-"`
}

func (Inbound) TableName() string {
	return "inbounds"
}
