package update_user

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, userID uint, request *user_dto.UpdateUserRequest) (*entities.User, error)
}
