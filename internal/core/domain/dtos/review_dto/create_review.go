package review_dto

type CreateReviewRequest struct {
	DoctorID      uint    `json:"doctor_id" validate:"required,min=1"`
	AppointmentID uint    `json:"appointment_id" validate:"required,min=1"`
	Rating        float64 `json:"rating" validate:"required,min=0.0,max=5.0"`
	Feedback      string  `json:"feedback" validate:"required,min=5,max=1024"`
}
