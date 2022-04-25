package models

import "github.com/google/uuid"

type Sale struct {
	BaseModel
	ItemID         uuid.UUID  `json:"item_id"`
	Item           Item       `json:"item"`
	Price          int        `json:"price"`
	CoinBuy        string     `json:"coin_buy"`
	CoinBuyAddress string     `json:"coin_buy_address"`
	Decimal        int        `json:"decimal"`
	BuyerID        *uuid.UUID `json:"buyer_id"`
}

func (Sale) TableName() string {
	return "sale"
}

type SaleReq struct {
	ItemID         uuid.UUID `json:"item_id"`
	Price          int       `json:"price"`
	CoinBuy        string    `json:"coin_buy"`
	CoinBuyAddress string    `json:"coin_buy_address"`
	Decimal        int       `json:"decimal"`
}

type QuerySaleReq struct {
	ItemID uuid.UUID `json:"item_id" form:"item_id"`
}

type QuerySaleRes struct {
	Data []Sale `json:"data"`
}
