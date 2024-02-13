package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mathews-r/golang/src/configs/rest_err"
)

const (
	JWT_SECRET = "JWT_SECRET"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.Email,
		"name":  ud.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerErr("Error to try generate jwt token")
	}

	return tokenString, nil
}
