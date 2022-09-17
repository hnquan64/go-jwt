package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(username string, email string) (tokenString string, err error) {
	// Phiên token: 1h
	expirationTime := time.Now().Add(1 * time.Hour)
	claim := &JWTClaim{
		UserName: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		return
	}
	claim, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claim")
	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
