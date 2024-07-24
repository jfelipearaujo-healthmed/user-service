package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/middlewares/role"
)

type token struct {
	signingKey string
}

func NewService(config *config.Config) TokenService {
	return &token{
		signingKey: config.TokenConfig.SignKey,
	}
}

func (t *token) CreateJwtToken(userID uint, role role.Role) (*Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  userID,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(t.signingKey))
	if err != nil {
		return nil, err
	}

	return NewBearer(tokenStr), nil
}
