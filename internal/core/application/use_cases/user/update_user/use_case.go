package update_user

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/update_user"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/fields"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
)

type useCase struct {
	database *persistence.Database
}

func NewUseCase(database *persistence.Database) contract.UseCase {
	return &useCase{
		database: database,
	}
}

func (uc *useCase) Execute(ctx context.Context, userID uint, newUser *entities.User) error {
	tx := uc.database.Instance.WithContext(ctx)

	fields := fields.GetNonEmptyFields(newUser, nil, nil)

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	result := tx.Model(&entities.User{}).
		Omit("id", "role", "doctor_id").
		Select("id", "full_name", "email", "password", "document_id", "phone").
		Where("id = ?", userID).
		Updates(fields)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
