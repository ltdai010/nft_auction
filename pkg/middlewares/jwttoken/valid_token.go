package jwttoken

import (
	// "ds/x/cryptosystem"
	"crypto/rsa"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtToken string

func (t JwtToken) String() string {
	return string(t)
}

func JwtTokenFrom(bearToken string) JwtToken {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return JwtToken(strArr[1])
	}
	return JwtToken(bearToken)
}

func (t JwtToken) TokenValid(publicKey *rsa.PublicKey) error {
	token, err := t.VerifyToken(publicKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (t JwtToken) VerifyToken(publicKey *rsa.PublicKey) (*jwt.Token, error) {

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

func (t JwtToken) ExtractTokenMetadata(publicKey *rsa.PublicKey) (*TokenInfo, error) {
	token, err := t.VerifyToken(publicKey)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		fmt.Println("mmm", TokenInfoFromClaims(claims))

		return TokenInfoFromClaims(claims), nil
	}
	return nil, err
}
