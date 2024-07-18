package get_user_by_id

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/get_user_by_id"
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

func (uc *useCase) Execute(ctx context.Context, id uint) (*entities.User, error) {
	tx := uc.database.Instance.WithContext(ctx)

	user := &entities.User{}

	result := tx.Preload("Doctor").First(user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return user, result.Error
}
