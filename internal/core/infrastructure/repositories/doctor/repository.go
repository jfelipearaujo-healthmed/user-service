package doctor_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	doctor_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/doctor"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type repository struct {
	dbService *persistence.DbService
}

func NewRepository(dbService *persistence.DbService) doctor_repository_contract.Repository {
	return &repository{
		dbService: dbService,
	}
}

func (rp *repository) GetByID(ctx context.Context, id uint) (*entities.Doctor, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	doctor := new(entities.Doctor)
	result := tx.First(doctor, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("doctor with id %d not found", id))
		}

		return nil, result.Error
	}

	return doctor, nil
}

func (rp *repository) GetByMedicalID(ctx context.Context, medicalID string, userIdToIgnore uint) (*entities.Doctor, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	doctor := new(entities.Doctor)
	result := tx.Where("medical_id = ? AND user_id != ?", medicalID, userIdToIgnore).First(doctor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("doctor with medical id %s not found", medicalID))
		}

		return nil, result.Error
	}

	return doctor, nil
}

func (rp *repository) Update(ctx context.Context, doctor *entities.Doctor) error {
	tx := rp.dbService.Instance.WithContext(ctx)

	if err := tx.Save(doctor).Error; err != nil {
		return err
	}

	return nil
}
