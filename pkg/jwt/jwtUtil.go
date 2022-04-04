package jwtUtils

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

//Our token struct
type DecodedToken struct {
	UserId string   `json:"userId"`
	Email  string   `json:"email"`
	Iat    int      `json:"iat"`
	Exp   int64 `json:"exp"`
	IsAdmin  string `json:"isAdmin"`
	// Iss    string   `json:"iss"`
}

//
func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}

func VerifyToken(token string, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken
}
