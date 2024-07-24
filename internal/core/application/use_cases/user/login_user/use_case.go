package login_user_uc

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	user_repository_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/repositories/user"
	login_user_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/login_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/token"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
)

type useCase struct {
	repository   user_repository_contract.Repository
	tokenService token.TokenService
	hasher       hasher.Hasher
}

func NewUseCase(
	repository user_repository_contract.Repository,
	tokenService token.TokenService,
	hasher hasher.Hasher,
) login_user_contract.UseCase {
	return &useCase{
		repository:   repository,
		tokenService: tokenService,
		hasher:       hasher,
	}
}

func (uc *useCase) Execute(ctx context.Context, request *user_dto.LoginUserRequest) (*token.Token, error) {
	if request.IsDoctorLogin() {
		slog.InfoContext(ctx, "login via doctor")

		user, err := uc.repository.GetByMedicalID(ctx, *request.MedicalID)
		if err != nil {
			slog.ErrorContext(ctx, "error getting user", "error", err)
			return nil, err
		}

		if !uc.hasher.ComparePassword(ctx, *request.Password, user.Password) {
			slog.ErrorContext(ctx, "invalid credentials, password does not match")
			return nil, app_error.New(http.StatusUnauthorized, "invalid credentials")
		}

		token, err := uc.tokenService.CreateJwtToken(user.ID, role.Doctor)
		if err != nil {
			slog.ErrorContext(ctx, "error creating token", "error", err)
			return nil, err
		}

		slog.InfoContext(ctx, "login via doctor successful")

		return token, nil
	}

	if request.IsPatientLogin() {
		slog.InfoContext(ctx, "login via patient")

		documentID := ""
		email := ""

		if request.DocumentID != nil {
			documentID = *request.DocumentID
		}
		if request.Email != nil {
			email = *request.Email
		}

		user, err := uc.repository.GetByDocumentIDOrEmail(ctx, documentID, email)
		if err != nil {
			slog.ErrorContext(ctx, "error getting user", "error", err)
			return nil, err
		}

		if !uc.hasher.ComparePassword(ctx, *request.Password, user.Password) {
			slog.ErrorContext(ctx, "invalid credentials, password does not match")
			return nil, app_error.New(http.StatusUnauthorized, "invalid credentials")
		}

		token, err := uc.tokenService.CreateJwtToken(user.ID, role.Patient)
		if err != nil {
			slog.ErrorContext(ctx, "error creating token", "error", err)
			return nil, err
		}

		slog.InfoContext(ctx, "login via patient successful")

		return token, nil
	}

	slog.ErrorContext(ctx, "invalid credentials, no login method found")

	return nil, app_error.New(http.StatusUnauthorized, "invalid credentials")
}
