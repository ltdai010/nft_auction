package jwttoken

import (
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenInfo struct {
	ID      uuid.UUID `json:"ID"`
	Pubkey  string    `json:"pubkey"`
	Expired int64     `json:"expired"`
}

func (t *TokenInfo) Claims() jwt.MapClaims {
	clams := jwt.MapClaims{}

	if t.ID != uuid.Nil {
		clams["id"] = t.ID
	} else {
		clams["id"] = uuid.New()
	}

	if t.Pubkey != "" {
		clams["userId"] = t.Pubkey
	}
	if t.Expired != 0 {
		clams["exp"] = t.Expired
	}
	return clams
}

func (t *TokenInfo) Create(privateKey *rsa.PrivateKey) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, t.Claims()).SignedString(privateKey)
}

func TokenInfoFromClaims(claims jwt.MapClaims) *TokenInfo {
	tokenInfo := &TokenInfo{}

	if cid, ok := claims["id"].(string); ok {
		if id, err := uuid.Parse(cid); err == nil {
			tokenInfo.ID = id
		}
	} else {
		tokenInfo.ID = uuid.New()
	}

	if cuid, ok := claims["userId"].(string); ok {
		tokenInfo.Pubkey = cuid
	}

	if exp, ok := claims["exp"].(int64); ok {
		tokenInfo.Expired = exp
	}

	return tokenInfo
}
