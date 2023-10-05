package token

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"
)

func GenerateToken(user_id uuid.UUID) (string, error) {

	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    user_id.String(),
		IssuedAt:  time.Now().Unix(),
	}
	secret := os.Getenv("API_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))

}
