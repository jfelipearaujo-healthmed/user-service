package create_user

import (
	"context"
	"errors"
	"net/http"

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

func (uc *useCase) Execute(ctx context.Context, user *entities.User) error {
	tx := uc.database.Instance.WithContext(ctx)

	existingUser := new(entities.User)
	if err := tx.Where("(document_id = ? OR email = ?) AND id != ?", user.DocumentID, user.Email, user.ID).First(existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if existingUser.ID != 0 {
		return app_error.New(http.StatusConflict, "user already exists")
	}

	if err := tx.Create(user).Error; err != nil {
		return err
	}

	return nil
}
