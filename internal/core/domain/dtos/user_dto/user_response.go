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

	Doctor    *DoctorResponse   `json:"doctor,omitempty"`
	Addresses []AddressResponse `json:"addresses,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DoctorResponse struct {
	MedicalID     string  `json:"medical_id"`
	Specialty     string  `json:"specialty"`
	Price         float64 `json:"price"`
	AvgRating     float64 `json:"avg_rating"`
	TotalPatients int     `json:"total_patients"`
}

type AddressResponse struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
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
		res.Doctor = &DoctorResponse{
			MedicalID:     user.Doctor.MedicalID,
			Specialty:     user.Doctor.Specialty,
			Price:         user.Doctor.Price,
			AvgRating:     user.Doctor.AvgRating,
			TotalPatients: user.Doctor.TotalPatients,
		}
	}

	res.Addresses = make([]AddressResponse, len(user.Addresses))

	for i := range user.Addresses {
		res.Addresses[i] = AddressResponse{
			Street:       user.Addresses[i].Street,
			Number:       user.Addresses[i].Number,
			Neighborhood: user.Addresses[i].Neighborhood,
			City:         user.Addresses[i].City,
			State:        user.Addresses[i].State,
			Zip:          user.Addresses[i].Zip,
		}
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
