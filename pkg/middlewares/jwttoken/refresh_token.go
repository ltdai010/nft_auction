package jwttoken

import (
	"crypto/rsa"
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type RefreshTokenInfo struct {
	ID      uuid.UUID `json:"id"`
	UserID  string    `json:"user_id"`
	Expired int64     `json:"expired"`
}

func (r *RefreshTokenInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (t *RefreshTokenInfo) Claims() jwt.MapClaims {
	clams := jwt.MapClaims{}

	if t.ID != uuid.Nil {
		clams["id"] = t.ID
	} else {
		clams["id"] = uuid.New()
	}

	if t.UserID != "" {
		clams["userId"] = t.UserID
	}

	if t.Expired != 0 {
		clams["exp"] = t.Expired
	}
	return clams
}

func (t *RefreshTokenInfo) Create(privateKey *rsa.PrivateKey) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, t.Claims()).SignedString(privateKey)
}

func RefreshTokenInfoFromClaims(claims jwt.MapClaims) *RefreshTokenInfo {
	tokenInfo := &RefreshTokenInfo{}

	if cid, ok := claims["id"].(string); ok {
		if id, err := uuid.Parse(cid); err == nil {
			tokenInfo.ID = id
		}
	} else {
		tokenInfo.ID = uuid.New()
	}

	if cuid, ok := claims["userId"].(string); ok {
		tokenInfo.UserID = cuid
	}

	if exp, ok := claims["exp"].(int64); ok {
		tokenInfo.Expired = exp
	}

	return tokenInfo
}
