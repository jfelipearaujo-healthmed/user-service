package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model

	UserID   string  `json:"user_id"`
	Rating   float64 `json:"rating"`
	Feedback string  `json:"feedback"`
}
