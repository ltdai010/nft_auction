package models

import "github.com/google/uuid"

type Order struct {
	BaseModel
	ItemID      uuid.UUID `json:"item_id"`
	UserID      uuid.UUID `json:"user_id"`
	User        User      `json:"user"`
	CoinName    string    `json:"coin_name"`
	CoinAddress string    `json:"coin_address"`
	Amount      int       `json:"amount"`
	Decimal     int       `json:"decimal"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderReq struct {
	CoinName    string `json:"coin_name"`
	CoinAddress string `json:"coin_address"`
	Amount      int    `json:"amount"`
	Decimal     int    `json:"decimal"`
}

type QueryOrderReq struct {
	ItemID string `json:"item_id"`
	Pagination
}

type QueryOrderRes struct {
	Data     []Order     `json:"data"`
	Metadata interface{} `json:"metadata"`
}
