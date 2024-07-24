package doctor_dto

type RatingDoctor struct {
	Rating float64 `json:"rating" validate:"required"`
}
