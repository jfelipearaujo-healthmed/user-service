package rating_doctor_uc

import (
	"context"

	doctor_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/doctor"
	rating_doctor_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/doctor/rating_doctor"
)

type useCase struct {
	repository doctor_repository_contract.Repository
}

func NewUseCase(repository doctor_repository_contract.Repository) rating_doctor_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, doctorID uint, rating float64) error {
	doctor, err := uc.repository.GetByID(ctx, doctorID)
	if err != nil {
		return err
	}

	doctor.AvgRating = rating
	doctor.TotalPatients++

	return uc.repository.Update(ctx, doctor)
}
