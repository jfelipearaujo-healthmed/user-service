package create_user

import (
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	create_user_uc "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/hasher"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/validator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase create_user_uc.UseCase
	hasher  hasher.Hasher
}

func NewHandler(useCase create_user_uc.UseCase, hasher hasher.Hasher) *handler {
	return &handler{
		useCase: useCase,
		hasher:  hasher,
	}
}

func (h *handler) Handle(c echo.Context) error {
	req := new(user_dto.CreateUserRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to read request body", err)
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	ctx := c.Request().Context()

	passHashed, err := h.hasher.HashPassword(req.Password)
	if err != nil {
		return http_response.BadRequest(c, "unable to hash the password", err)
	}

	user := &entities.User{
		FullName:   req.FullName,
		Email:      req.Email,
		Password:   passHashed,
		DocumentID: req.DocumentID,
		Phone:      req.Phone,
		Role:       req.Role,
	}

	if user.IsDoctor() {
		doctor := &entities.Doctor{
			MedicalID: req.DoctorMedicalID,
			Specialty: req.DoctorSpecialty,
			Price:     req.DoctorPrice,
		}

		if err := validator.Validate(doctor); err != nil {
			return http_response.UnprocessableEntity(c, "invalid request body", err)
		}

		user.Doctor = doctor
	}

	if err := h.useCase.Execute(ctx, user); err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, user_dto.MapFromDomain(user))
}
