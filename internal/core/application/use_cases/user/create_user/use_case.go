package create_user

import (
	"context"
	"errors"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
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

func (uc *useCase) Execute(ctx context.Context, request *user_dto.CreateUserRequest) (*entities.User, error) {
	tx := uc.database.Instance.WithContext(ctx)

	existingUser := new(entities.User)
	if err := tx.Where("document_id = ? OR email = ?", request.DocumentID, request.Email).First(existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if existingUser.ID != 0 {
		return nil, app_error.New(http.StatusConflict, "user already exists")
	}

	user := &entities.User{
		FullName:   request.FullName,
		Email:      request.Email,
		Password:   request.Password,
		DocumentID: request.DocumentID,
		Phone:      request.Phone,
		Role:       request.Role,
	}

	if user.IsDoctor() {
		user.Doctor = &entities.Doctor{
			MedicalID: *request.DoctorMedicalID,
			Specialty: *request.DoctorSpecialty,
			Price:     *request.DoctorPrice,
		}
	}

	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
