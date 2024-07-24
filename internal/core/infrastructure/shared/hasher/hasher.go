package hasher

import (
	"context"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
}

func NewHasher() Hasher {
	return &service{}
}

func (h *service) HashPassword(ctx context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.ErrorContext(ctx, "error hashing password", "error", err)
		return "", err
	}
	return string(bytes), nil
}

func (h *service) ComparePassword(ctx context.Context, password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		slog.ErrorContext(ctx, "error comparing password", "error", err)
		return false
	}
	return true
}
