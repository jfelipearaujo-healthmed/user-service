package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DocumentID string `json:"document_id"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`

	DoctorID *string  `json:"doctor_id"`
	Doctor   *Doctor  `json:"doctor"`
	Reviews  []Review `json:"reviews"`
}
