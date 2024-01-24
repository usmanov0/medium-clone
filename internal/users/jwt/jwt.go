package jwt

import (
	"example.com/my-medium-clone/internal/errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = []byte("secret-key")

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(req string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"expiredAt": expirationTime.Unix(),
		"issuedAt":  time.Now().Unix(),
		"red":       req,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.ErrCreationToken
	}
	return signedToken, nil
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
