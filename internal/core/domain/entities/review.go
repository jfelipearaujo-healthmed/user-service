package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model

	DoctorID  uint    `json:"doctor_id"`
	PatientID uint    `json:"patient_id"`
	Rating    float64 `json:"rating"`
	Feedback  string  `json:"feedback"`

	Doctor  Doctor `gorm:"foreignKey:DoctorID"`
	Patient User   `gorm:"foreignKey:PatientID"`
}
