package jwttoken

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type RefreshTokenInfo struct {
	ID      uuid.UUID `json:"id"`
	UserID  string    `json:"user_id"`
	Pubkey  string    `json:"pubkey"`
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

	if t.Pubkey != "" {
		clams["pubkey"] = t.Pubkey
	}

	if t.Expired != 0 {
		clams["exp"] = fmt.Sprint(t.Expired)
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

	if pubkey, ok := claims["pubkey"].(string); ok {
		tokenInfo.Pubkey = pubkey
	}

	if exp, err := strconv.ParseInt(fmt.Sprint(claims["exp"]), 10, 64); err == nil {
		tokenInfo.Expired = exp
	}

	return tokenInfo
}
