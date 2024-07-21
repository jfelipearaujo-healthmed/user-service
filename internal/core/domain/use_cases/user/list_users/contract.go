package list_users_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
)

type Filter struct {
	DocumentID *string `query:"document_id" json:"document_id" validate:"omitempty,cpfcnpj"`
	Email      *string `query:"email" json:"email" validate:"omitempty,email"`
	FullName   *string `query:"full_name" json:"full_name" validate:"omitempty,min=5,max=255"`
	Phone      *string `query:"phone" json:"phone" validate:"omitempty,min=5,max=255"`

	MedicalID *string  `query:"medical_id" json:"medical_id" validate:"omitempty,min=2,max=255"`
	Specialty *string  `query:"specialty" json:"specialty" validate:"omitempty,min=5,max=255"`
	AvgRating *float64 `query:"avg_rating" json:"avg_rating" validate:"omitempty,min=1,max=5"`

	RoleFilter role.Role
}

func (f *Filter) IsEmpty() bool {
	return f.DocumentID == nil &&
		f.Email == nil &&
		f.FullName == nil &&
		f.Phone == nil &&
		f.MedicalID == nil &&
		f.Specialty == nil &&
		f.AvgRating == nil
}

type UseCase interface {
	Execute(ctx context.Context, filter *Filter) ([]entities.User, error)
}
