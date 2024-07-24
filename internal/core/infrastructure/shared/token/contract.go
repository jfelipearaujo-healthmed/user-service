package token

import "github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"

type TokenService interface {
	CreateJwtToken(userID uint, role role.Role) (*Token, error)
}

type Token struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

func NewBearer(token string) *Token {
	return &Token{
		Type:  "Bearer",
		Token: token,
	}
}
