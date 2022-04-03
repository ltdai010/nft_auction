package models

type LoginTokenResponse struct {
	Token        string `json:"token"`
	ExpiredAt    int64  `json:"expired_at"`
	RefreshToken string `json:"refresh_token"`
}

type NewTokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
