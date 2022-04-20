package models

import "github.com/google/uuid"

type Item struct {
	BaseModel
	ItemID       string     `json:"item_id" gorm:"index"`
	Owner        string     `json:"owner"`
	CollectionID uuid.UUID  `json:"collection_id"`
	Collection   Collection `json:"collection"`
}

func (user *Item) TableName() string {
	return "items"
}

type QueryItemReq struct {
	Name         string `json:"name"`
	CollectionID string `json:"collection_id"`
	OwnerID      string `json:"owner_id"`
	Pagination
}

type QueryItemRes struct {
	Data     []Collection `json:"data"`
	Metadata interface{}  `json:"metadata"`
}
