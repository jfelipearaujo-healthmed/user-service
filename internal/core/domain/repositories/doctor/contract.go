package doctor_repository_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type Repository interface {
	GetByID(ctx context.Context, id uint) (*entities.Doctor, error)
	GetByMedicalID(ctx context.Context, medicalID string, userIdToIgnore uint) (*entities.Doctor, error)
	Update(ctx context.Context, doctor *entities.Doctor) error
}
