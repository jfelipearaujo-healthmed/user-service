package delete_user

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/delete_user"
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

func (uc *useCase) Execute(ctx context.Context, id uint) error {
	tx := uc.database.Instance.WithContext(ctx)

	result := tx.Delete(&entities.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
