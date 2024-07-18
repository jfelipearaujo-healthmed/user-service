package delete_user_contract

import "context"

type UseCase interface {
	Execute(ctx context.Context, id uint) error
}
