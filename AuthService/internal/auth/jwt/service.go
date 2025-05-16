package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey []byte
}

func NewJWTService(secretKey []byte) jwtService {
	return jwtService{
		secretKey: secretKey,
	}
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Role     string `json:"role"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func createJWTClaims(role, email, nickname string) *jwtClaims {
	return &jwtClaims{
		Role:     role,
		Email:    email,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
}

func (js *jwtService) CreateAccessToken(role, email, nickname string) (string, error) {
	claims := createJWTClaims(role, email, nickname)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(js.secretKey)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrSignToken, err)
	}

	return tokenString, nil
}

func (jw *jwtClaims) VerifyToken(token string) {
}
