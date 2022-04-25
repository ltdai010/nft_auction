package models

import (
	"encoding/json"
	"time"
)

type User struct {
	BaseModel
	LastLogin time.Time        `json:"last_login" gorm:"autoCreateTime" sql:"default:CURRENT_TIMESTAMP"`
	Pubkey    string           `json:"pubkey" gorm:"uniqueIndex"`
	Address   string           `json:"address"`
	Name      string           `json:"name"`
	Metadata  *json.RawMessage `json:"metadata" swaggertype:"object"`
}

type UsersLogin struct {
	Data  User               `json:"data"`
	Token LoginTokenResponse `json:"token"`
}

type UserLoginRequest struct {
	Signature string `json:"signature"`
}

type UserRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (user *User) TableName() string {
	return "users"
}
