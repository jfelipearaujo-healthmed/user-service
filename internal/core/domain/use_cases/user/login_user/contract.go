package login_user_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/domain/dtos/user_dto"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/token"
)

type UseCase interface {
	Execute(ctx context.Context, request *user_dto.LoginUserRequest) (*token.Token, error)
}
