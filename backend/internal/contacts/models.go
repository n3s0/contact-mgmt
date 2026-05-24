package contacts

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Phone string `json:"phone"`
	Image string `json:"image"`
}

func (Contact) TableName() string {
	return "contacts"
}
