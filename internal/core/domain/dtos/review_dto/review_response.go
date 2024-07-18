package review_dto

import (
	"time"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type ReviewResponse struct {
	ID uint `json:"id"`

	DoctorID      uint    `json:"doctor_id"`
	PatientID     uint    `json:"patient_id"`
	AppointmentID uint    `json:"appointment_id"`
	Rating        float64 `json:"rating"`
	Feedback      string  `json:"feedback"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapFromDomain(review *entities.Review) *ReviewResponse {
	return &ReviewResponse{
		ID: review.ID,

		DoctorID:      review.DoctorID,
		PatientID:     review.PatientID,
		AppointmentID: review.AppointmentID,
		Rating:        review.Rating,
		Feedback:      review.Feedback,

		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}

func MapFromDomainSlice(reviews []entities.Review) []*ReviewResponse {
	reviewResponses := make([]*ReviewResponse, len(reviews))

	for i, review := range reviews {
		reviewResponses[i] = MapFromDomain(&review)
	}

	return reviewResponses
}
