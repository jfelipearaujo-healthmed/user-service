package delete_user

import "context"

type UseCase interface {
	Execute(ctx context.Context, id uint) error
}
