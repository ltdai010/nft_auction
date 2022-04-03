package models

import (
	"encoding/json"
	"time"
)

type Users struct {
	BaseModel
	LastLogin time.Time        `json:"last_login" gorm:"autoCreateTime" sql:"default:CURRENT_TIMESTAMP"`
	Pubkey    string           `json:"pubkey" gorm:"uniqueIndex"`
	Address   string           `json:"address"`
	Name      string           `json:"name"`
	Metadata  *json.RawMessage `json:"metadata"`
}

type UsersLogin struct {
	Data  Users              `json:"data"`
	Token LoginTokenResponse `json:"token"`
}

type UserLoginRequest struct {
	Signature string `json:"signature"`
}

func (user *Users) TableName() string {
	return "users"
}
