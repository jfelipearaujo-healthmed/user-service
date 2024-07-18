package update_user

import (
	"context"
	"errors"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
	"gorm.io/gorm"
)

type useCase struct {
	database *persistence.DbService
}

func NewUseCase(database *persistence.DbService) contract.UseCase {
	return &useCase{
		database: database,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint, request *user_dto.UpdateUserRequest) (*entities.User, error) {
	tx := uc.database.Instance.WithContext(ctx)

	user := &entities.User{}

	result := tx.Preload("Doctor").First(user, userID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, "user not found")
		}

		return nil, result.Error
	}

	if request.FullName != nil && user.FullName != *request.FullName {
		user.FullName = *request.FullName
	}
	if request.Email != nil && user.Email != *request.Email {
		user.Email = *request.Email
	}
	if request.Password != nil && user.Password != *request.Password {
		user.Password = *request.Password
	}
	if request.DocumentID != nil && user.DocumentID != *request.DocumentID {
		user.DocumentID = *request.DocumentID
	}
	if request.Phone != nil && user.Phone != *request.Phone {
		user.Phone = *request.Phone
	}

	if user.IsDoctor() {
		if request.DoctorMedicalID != nil && user.Doctor.MedicalID != *request.DoctorMedicalID {
			user.Doctor.MedicalID = *request.DoctorMedicalID
		}
		if request.DoctorSpecialty != nil && user.Doctor.Specialty != *request.DoctorSpecialty {
			user.Doctor.Specialty = *request.DoctorSpecialty
		}
		if request.DoctorPrice != nil && user.Doctor.Price != *request.DoctorPrice {
			user.Doctor.Price = *request.DoctorPrice
		}
	}

	if request.Email != nil || request.DocumentID != nil {
		existingUser := new(entities.User)
		if err := tx.Where("(document_id = ? OR email = ?) AND id != ?", request.DocumentID, request.Email, userID).First(existingUser).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		if existingUser.ID != 0 {
			return nil, app_error.New(http.StatusConflict, "e-mail or document id in use")
		}
	}

	if request.DoctorMedicalID != nil {
		existingDoctor := new(entities.Doctor)
		if err := tx.Where("medical_id = ? AND user_id != ?", *request.DoctorMedicalID, userID).First(existingDoctor).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}

		if existingDoctor.ID != 0 {
			return nil, app_error.New(http.StatusConflict, "medical id in use")
		}
	}

	result = tx.Model(user).Save(user)

	if err := result.Error; err != nil {
		return nil, err
	}

	if user.IsDoctor() {
		result = tx.Model(user.Doctor).Save(user.Doctor)
		if err := result.Error; err != nil {
			return nil, err
		}
	}

	return user, nil
}
