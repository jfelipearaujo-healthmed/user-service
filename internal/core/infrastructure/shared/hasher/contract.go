package hasher

import "context"

type Hasher interface {
	HashPassword(ctx context.Context, password string) (string, error)
	ComparePassword(ctx context.Context, password string, hashedPassword string) bool
}
