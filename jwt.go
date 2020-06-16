package common

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTSecret jwt token secret, from ENV kubernetes
var JWTSecret = []byte(os.Getenv("JWT_KEY"))

//TokenParam param for generate JWT token
type TokenParam struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
	Role      int    `json:"role"`
}

//ClaimsToken jwt token claims struct
type ClaimsToken struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
	Role      int    `json:"role"`
	Expired   int64  `json:"exp"`
	jwt.StandardClaims
}

//GenerateJWT generate JWT Token using Mehod_HS256
//expire 24 hours
func GenerateJWT(claimParam TokenParam) (token string, err error) {
	claims := &ClaimsToken{
		UserID:    claimParam.UserID,
		Email:     claimParam.Email,
		Signature: claimParam.Signature,
		Role:      claimParam.Role,
		Expired:   time.Now().Add(time.Hour * 24).Unix(), // 24 hours expired. manual implement refresh token
	}

	signJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = signJwt.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

//DecodeJWT decode jwt to be struct
func DecodeJWT(token string) (*ClaimsToken, error) {
	claims := &ClaimsToken{}
	tkn, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, ErrUnAuthorized
		}
		return claims, err
	}

	if !tkn.Valid {
		return claims, ErrUnAuthorized
	}

	return claims, nil
}
