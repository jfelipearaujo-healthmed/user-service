package user_dto

import (
	"time"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type UserResponse struct {
	ID uint `json:"id"`

	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	DocumentID string `json:"document_id"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`

	DoctorMedicalID string  `json:"medical_id,omitempty"`
	DoctorSpecialty string  `json:"specialty,omitempty"`
	DoctorPrice     float64 `json:"price,omitempty"`
	AvgRating       float64 `json:"avg_rating,omitempty"`
	TotalPatients   int     `json:"total_patients,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapFromDomain(user *entities.User) *UserResponse {
	res := &UserResponse{
		ID: user.ID,

		FullName:   user.FullName,
		Email:      user.Email,
		DocumentID: user.DocumentID,
		Phone:      user.Phone,
		Role:       user.Role,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Doctor != nil {
		res.DoctorMedicalID = user.Doctor.MedicalID
		res.DoctorSpecialty = user.Doctor.Specialty
		res.DoctorPrice = user.Doctor.Price
		res.AvgRating = user.Doctor.AvgRating
		res.TotalPatients = user.Doctor.TotalPatients
	}

	return res
}

func MapFromSlice(users []entities.User) []*UserResponse {
	res := make([]*UserResponse, len(users))

	for i := range users {
		res[i] = MapFromDomain(&users[i])
	}

	return res
}
