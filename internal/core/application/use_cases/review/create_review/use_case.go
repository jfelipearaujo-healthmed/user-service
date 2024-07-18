package create_review

import (
	"context"
	"errors"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/review_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/review/create_review"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type useCase struct {
	dbService *persistence.DbService
}

func NewUseCase(dbService *persistence.DbService) contract.UseCase {
	return &useCase{
		dbService: dbService,
	}
}

func (uc *useCase) Execute(ctx context.Context, patientID uint, request *review_dto.CreateReviewRequest) (*entities.Review, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	patient, err := uc.getPatient(ctx, patientID)
	if err != nil {
		return nil, err
	}

	if patient.IsDoctor() {
		return nil, app_error.New(http.StatusBadRequest, "user not allowed to add a review, must be a patient")
	}

	if patient.ID == 0 {
		return nil, app_error.New(http.StatusNotFound, "patient not found")
	}

	existingReview, err := uc.getReview(ctx, patientID, request)
	if err != nil {
		return nil, err
	}

	if existingReview.ID != 0 {
		return nil, app_error.New(http.StatusConflict, "review already exists")
	}

	doctor, err := uc.getDoctor(ctx, request.DoctorID)
	if err != nil {
		return nil, err
	}

	if doctor.ID == 0 {
		return nil, app_error.New(http.StatusNotFound, "doctor not found")
	}

	review := &entities.Review{
		DoctorID:      request.DoctorID,
		PatientID:     patientID,
		AppointmentID: request.AppointmentID,
		Rating:        request.Rating,
		Feedback:      request.Feedback,
	}

	if err := tx.Create(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (uc *useCase) getReview(ctx context.Context, patientID uint, request *review_dto.CreateReviewRequest) (*entities.Review, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	existingReview := new(entities.Review)
	if err := tx.Where("patient_id = ? AND doctor_id = ? AND appointment_id = ?",
		patientID,
		request.DoctorID,
		request.AppointmentID).
		First(existingReview).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return existingReview, nil
}

func (uc *useCase) getPatient(ctx context.Context, patientID uint) (*entities.User, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	user := new(entities.User)
	if err := tx.Where("id = ?", patientID).First(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return user, nil
}

func (uc *useCase) getDoctor(ctx context.Context, doctorID uint) (*entities.Doctor, error) {
	tx := uc.dbService.Instance.WithContext(ctx)

	doctor := new(entities.Doctor)
	if err := tx.Where("id = ?", doctorID).First(doctor).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return doctor, nil
}
