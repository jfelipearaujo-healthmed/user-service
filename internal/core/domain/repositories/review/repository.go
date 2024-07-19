package review_repository_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type Repository interface {
	GetByID(ctx context.Context, reviewID, userID uint, isDoctor bool) (*entities.Review, error)
	GetByIDs(ctx context.Context, patientID, doctorID, appointmentID uint) (*entities.Review, error)
	GetByPatientID(ctx context.Context, patientID uint) (*entities.Review, error)
	GetByDoctorID(ctx context.Context, doctorID uint) (*entities.Review, error)
	List(ctx context.Context, userID uint, isDoctor bool) ([]entities.Review, error)
	Create(ctx context.Context, review *entities.Review) (*entities.Review, error)
	Update(ctx context.Context, review *entities.Review) (*entities.Review, error)
	Delete(ctx context.Context, id uint) error
}
