package list_users_uc

import (
	"context"
	"fmt"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	list_users_contract "github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/use_cases/user/list_users"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/fields"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/persistence"
)

type useCase struct {
	database *persistence.DbService
}

func NewUseCase(database *persistence.DbService) list_users_contract.UseCase {
	return &useCase{
		database: database,
	}
}

func (uc *useCase) Execute(ctx context.Context, filter *list_users_contract.Filter) ([]*entities.User, error) {
	tx := uc.database.Instance.WithContext(ctx)

	users := []*entities.User{}

	fields := fields.GetNonEmptyFields(filter, &fields.ANY_CHAR, &fields.ANY_CHAR)

	query := tx

	for field, value := range fields {
		query = query.Where(fmt.Sprintf("%s LIKE ?", field), value)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
