package utils

import (
	"time"

	"github.com/Abdulqudri/myapi/configs"
	"github.com/golang-jwt/jwt/v5"
)

//var jwtKey = configs.GetTokenSecret()

func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),  // Token valid for 15 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := configs.GetTokenSecret()

	return token.SignedString([]byte(jwtKey))

}
func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // Token valid for 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := configs.GetTokenSecret()

	return token.SignedString([]byte(jwtKey))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	jwtKey := configs.GetTokenSecret()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
