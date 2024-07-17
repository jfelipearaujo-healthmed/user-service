package create_user

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/create_user"
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

func (uc *useCase) Execute(ctx context.Context, user *entities.User) error {
	tx := uc.database.Instance.WithContext(ctx)

	if err := tx.Create(user).Error; err != nil {
		return err
	}

	return nil
}
