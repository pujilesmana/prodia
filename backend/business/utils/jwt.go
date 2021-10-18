package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaimsUser struct {
	ID    string
	Admin bool
	jwt.StandardClaims
}

func GenerateToken(id string) (string, error) {
	var jwtKey = os.Getenv("JWT_KEY")
	claims := JwtCustomClaimsUser{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}
