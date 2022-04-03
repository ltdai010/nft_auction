package jwttoken

import (
	// "ds/x/cryptosystem"
	"crypto/rsa"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtRefreshToken string

func (t JwtRefreshToken) String() string {
	return string(t)
}

func JwtRefreshTokenFrom(refreshToken string) JwtRefreshToken {
	return JwtRefreshToken(refreshToken)
}

func (t JwtRefreshToken) TokenValid(publicKey *rsa.PublicKey) error {
	token, err := t.VerifyToken(publicKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (t JwtRefreshToken) VerifyToken(publicKey *rsa.PublicKey) (*jwt.Token, error) {

	token, err := jwt.Parse(t.String(), func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (t JwtRefreshToken) ExtractTokenMetadata(publicKey *rsa.PublicKey) (*RefreshTokenInfo, error) {
	token, err := t.VerifyToken(publicKey)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return RefreshTokenInfoFromClaims(claims), nil
	}
	return nil, err
}
