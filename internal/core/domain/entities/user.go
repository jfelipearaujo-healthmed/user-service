package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FullName   string `json:"full_name"`
	Email      string `json:"email" gorm:"uniqueIndex"`
	Password   string `json:"password"`
	DocumentID string `json:"document_id" gorm:"uniqueIndex"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`

	Doctor *Doctor `gorm:"foreignKey:UserID"`
}

func (u *User) IsDoctor() bool {
	return u.Role == "doctor"
}
