package jwtUtils

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

//Our token struct
// Reserved claims: https://tools.ietf.org/html/rfc7519#section-4.1
// iss (issuer), *exp (expiration time), sub (subject), aud (audience)
type DecodedJWTToken struct {
	UserId  string `json:"userId"`
	Email   string `json:"email"`
	Iat     int    `json:"iat"` //issued at  *optional
	Exp     int64  `json:"exp"` //expiration time *Must be used
	IsAdmin bool `json:"isAdmin"`
	IsItAccesToken bool `json:"isItAccesToken"`
	// Iss    string   `json:"iss"`
}

/*
https://auth0.com/learn/json-web-tokens/
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)

  header.payload.signature

*/
func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}

func VerifyToken(token string, secret string) *DecodedJWTToken {
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

	var decodedToken DecodedJWTToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken
}
