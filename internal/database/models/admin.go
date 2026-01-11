package models

type Admin struct {
	Name     string `gorm:"type:text;unique;not null"`
	UserUUID string `gorm:"column:user_uuid;type:varchar(36);not null;unique;index;constraint:OnDelete:CASCADE"`

	User User `gorm:"foreignKey:UserUUID;references:UUID"`
}

func (Admin) TableName() string {
	return "admins"
}
