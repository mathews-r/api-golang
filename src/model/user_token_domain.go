package model

import (
	"os"
	"strings"
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

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	token, err := jwt.Parse(RemoveBearerFromToken(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestErr("Invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedErr("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedErr("Invalid token")
	}

	return &userDomain{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
	}, nil
}

func RemoveBearerFromToken(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}
	// return token[7:]
	return token
}
