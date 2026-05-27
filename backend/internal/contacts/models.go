package contacts

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName string
	LastName string
	Email string `gorm:"unique"`
	Phone string `gorm:"unique"`
	Image string
}

func (Contact) TableName() string {
	return "contacts"
}
