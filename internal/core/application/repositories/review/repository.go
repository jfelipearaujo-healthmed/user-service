package review_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	review_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/review"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type repository struct {
	dbService *persistence.DbService
}

func NewRepository(dbService *persistence.DbService) review_repository_contract.Repository {
	return &repository{
		dbService: dbService,
	}
}

func (rp *repository) GetByID(ctx context.Context, reviewID, userID uint, isDoctor bool) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	review := new(entities.Review)

	query := "patient_id = ? AND id = ?"

	if isDoctor {
		query = "doctor_id = ? AND id = ?"
	}

	result := tx.
		Preload("Doctor").
		Preload("Patient").
		Where(query, userID, reviewID).
		First(review)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("review with id %d not found", reviewID))
		}

		return nil, result.Error
	}

	return review, nil
}

func (rp *repository) GetByIDs(ctx context.Context, patientID uint, doctorID uint, appointmentID uint) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	review := new(entities.Review)
	result := tx.
		Preload("Doctor").
		Preload("Patient").
		Where("patient_id = ? AND doctor_id = ? AND appointment_id = ?", patientID, doctorID, appointmentID).
		First(review)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("review with patient id %d, doctor id %d and appointment id %d not found", patientID, doctorID, appointmentID))
		}

		return nil, result.Error
	}

	return review, nil
}

func (rp *repository) GetByPatientID(ctx context.Context, patientID uint) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	review := new(entities.Review)
	result := tx.
		Preload("Doctor").
		Preload("Patient").
		Where("patient_id = ?", patientID).
		First(review)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("review with patient id %d not found", patientID))
		}

		return nil, result.Error
	}

	return review, nil
}

func (rp *repository) GetByDoctorID(ctx context.Context, doctorID uint) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	review := new(entities.Review)
	result := tx.
		Preload("Doctor").
		Preload("Patient").
		Where("doctor_id = ?", doctorID).
		First(review)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("review with doctor id %d not found", doctorID))
		}

		return nil, result.Error
	}

	return review, nil
}

func (rp *repository) List(ctx context.Context, userID uint, isDoctor bool) ([]entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	reviews := new([]entities.Review)

	query := "patient_id = ?"

	if isDoctor {
		query = "doctor_id = ?"
	}

	result := tx.
		Preload("Doctor").
		Preload("Patient").
		Where(query, userID).
		Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return *reviews, nil
}

func (rp *repository) Create(ctx context.Context, review *entities.Review) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	if err := tx.Create(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (rp *repository) Update(ctx context.Context, review *entities.Review) (*entities.Review, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	result := tx.Model(review).Save(review)

	if err := result.Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (rp *repository) Delete(ctx context.Context, id uint) error {
	tx := rp.dbService.Instance.WithContext(ctx)

	result := tx.Delete(&entities.Review{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
