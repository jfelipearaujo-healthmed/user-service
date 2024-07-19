package update_user_uc

import (
	"context"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	doctor_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/doctor"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	update_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
)

type useCase struct {
	userRepository   user_repository_contract.Repository
	doctorRepository doctor_repository_contract.Repository
}

func NewUseCase(
	userRepository user_repository_contract.Repository,
	doctorRepository doctor_repository_contract.Repository,
) update_user_contract.UseCase {
	return &useCase{
		userRepository:   userRepository,
		doctorRepository: doctorRepository,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint, request *user_dto.UpdateUserRequest) (*entities.User, error) {
	user, err := uc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
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
		var documentId, email string

		if request.DocumentID != nil {
			documentId = *request.DocumentID
		}
		if request.Email != nil {
			email = *request.Email
		}

		existingUser, err := uc.userRepository.GetByDocumentIDOrEmail(ctx, documentId, email)
		if err != nil && !app_error.IsAppError(err) {
			return nil, err
		}

		if existingUser != nil {
			return nil, app_error.New(http.StatusConflict, "e-mail or document id in use")
		}
	}

	if request.DoctorMedicalID != nil {
		existingDoctor, err := uc.doctorRepository.GetByMedicalID(ctx, *request.DoctorMedicalID, userID)
		if err != nil && !app_error.IsAppError(err) {
			return nil, err
		}

		if existingDoctor != nil {
			return nil, app_error.New(http.StatusConflict, "medical id in use")
		}
	}

	user, err = uc.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
