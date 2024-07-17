package update_user

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, userID uint, newUser *entities.User) error
}
