package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model

	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`

	UserID uint `json:"user_id"`
}
