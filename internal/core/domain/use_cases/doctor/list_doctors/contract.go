package list_doctors_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type Filter struct {
	FullName  *string  `json:"full_name"`
	MedicalID *string  `json:"medical_id"`
	Specialty *string  `json:"specialty"`
	AvgRating *float64 `json:"avg_rating"`
}

type UseCase interface {
	Execute(ctx context.Context, filter *Filter) ([]entities.User, error)
}
