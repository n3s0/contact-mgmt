package addresses

import "gorm.io/gorm"

type Address struct {
	Street string
	CityID uint
	City City
	StateID uint
	State State
	PostalCode uint
	CountryID uint
	Country Country
}

type City struct {
	gorm.Model
	Name string
}

type State struct {
	gorm.Model
	Name string `gorm:"unique"`
	Abbreviation string `gorm:"unique"`
}

type Country struct {
	gorm.Model
	Name string `gorm:"unique"`
	Abbreviation string `gorm:"unique"`
}
