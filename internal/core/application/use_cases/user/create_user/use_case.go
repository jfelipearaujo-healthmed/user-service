package create_user_uc

import (
	"context"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	create_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
)

type useCase struct {
	repository user_repository_contract.Repository
}

func NewUseCase(repository user_repository_contract.Repository) create_user_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, request *user_dto.CreateUserRequest) (*entities.User, error) {
	existingUser, err := uc.repository.GetByDocumentIDOrEmail(ctx, request.DocumentID, request.Email)
	if err != nil && !app_error.IsAppError(err) {
		return nil, err
	}

	if existingUser != nil {
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

	return uc.repository.Create(ctx, user)
}
