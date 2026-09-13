package jwt

import "github.com/golang-jwt/jwt/v5"

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}