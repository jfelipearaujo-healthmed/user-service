package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	HashPassword(password string) (string, error)
}

type hasher struct {
}

func NewHasher() Hasher {
	return &hasher{}
}

func (h *hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
