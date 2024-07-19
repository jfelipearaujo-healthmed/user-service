package create_review_uc

import (
	"context"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	review_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/review"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/create_review"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
)

type useCase struct {
	reviewRepository review_repository_contract.Repository
	userRepository   user_repository_contract.Repository
}

func NewUseCase(
	reviewRepository review_repository_contract.Repository,
	userRepository user_repository_contract.Repository,
) contract.UseCase {
	return &useCase{
		reviewRepository: reviewRepository,
		userRepository:   userRepository,
	}
}

func (uc *useCase) Execute(ctx context.Context, patientID uint, request *review_dto.CreateReviewRequest) (*entities.Review, error) {
	user, err := uc.userRepository.GetByID(ctx, patientID)
	if err != nil {
		return nil, err
	}

	if user.IsDoctor() {
		return nil, app_error.New(http.StatusBadRequest, "user not allowed to add a review, must be a patient")
	}

	existingReview, err := uc.reviewRepository.GetByIDs(ctx, user.ID, request.DoctorID, request.AppointmentID)
	if err != nil && !app_error.IsAppError(err) {
		return nil, err
	}

	if existingReview != nil {
		return nil, app_error.New(http.StatusConflict, "review already exists")
	}

	doctor, err := uc.userRepository.GetByID(ctx, request.DoctorID)
	if err != nil {
		return nil, err
	}

	review := &entities.Review{
		DoctorID:      doctor.ID,
		PatientID:     user.ID,
		AppointmentID: request.AppointmentID,
		Rating:        request.Rating,
		Feedback:      request.Feedback,
	}

	return uc.reviewRepository.Create(ctx, review)
}
