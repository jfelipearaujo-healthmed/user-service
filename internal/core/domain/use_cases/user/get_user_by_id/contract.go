package get_user_by_id_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/entities"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
)

type UseCase interface {
	Execute(ctx context.Context, userID uint, roleFilter role.Role) (*entities.User, error)
}
