package addresses

import (
	"gorm.io/gorm"
	"contact-mgmt-backend/internal/contacts"
)

type Address struct {
	gorm.Model
	Street string
	CityID uint `gorm:"uniqueIndex;not null"`
	StateID uint `gorm"uniqueIndex;not null"`
	CountyID uint `gorm:"uniqueIndex"`
	PostalCode uint
	CountryID uint `gorm:"uniqueIndex"`
	ContactID *contacts.Contact `gorm:"foreignKey:ContactID"`
}

type City struct {
	gorm.Model
	Name string
	Address []Address `gorm:"foreignKey:AddressID"`
}

type State struct {
	gorm.Model
	Name string `gorm:"unique"`
	Abbreviation string `gorm:"unique"`
	Address []Address `gorm:"foreignKey:AddressID"`
}

type Country struct {
	gorm.Model
	Name string `gorm:"unique"`
	Abbreviation string `gorm:"unique"`
	Address []Address `gorm:"foreignKey:AddressID"`
}

func (Address) TableName() string {
	return "addresses"
}

func (City) TableName() string {
	return "cities"
}

func (State) TableName() string {
	return "states"
}

func (Country) TableName() string {
	return "countries"
}
