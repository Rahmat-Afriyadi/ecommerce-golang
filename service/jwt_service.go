package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaim struct {
	UserID string
	*jwt.StandardClaims
}

type JwtServiceImpl struct {
	issuer    string
	secretKey string
}

func NewJWTServiceImpl(secret string) *JwtServiceImpl {
	return &JwtServiceImpl{
		secretKey: secret,
	}
}

func (service *JwtServiceImpl) GenerateToken(userID string) string {
	claims := &JwtCustomClaim{
		userID,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    service.secretKey,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

