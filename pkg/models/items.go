package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Item struct {
	BaseModel
	ItemID       string          `json:"item_id" gorm:"index"`
	Owner        string          `json:"owner"`
	CollectionID uuid.UUID       `json:"collection_id"`
	Collection   Collection      `json:"collection"`
	Metadata     json.RawMessage `json:"metadata" swaggertype:"object"`
	Category     string          `json:"category"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
}

func (i *Item) TableName() string {
	return "items"
}

type ItemLike struct {
	BaseModel
	ItemID uuid.UUID `json:"item_id" gorm:"uniqueIndex:idx_item_user"`
	Item   Item      `json:"item"`
	UserID uuid.UUID `json:"user_id" gorm:"uniqueIndex:idx_item_user"`
	User   User      `json:"user"`
}

func (il *ItemLike) TableName() string {
	return "item_like"
}

type QueryItemReq struct {
	Name         string `json:"name"`
	CollectionID string `json:"collection_id"`
	OwnerID      string `json:"owner_id"`
	LikedBy      string `json:"liked_by"`
	CreatorID    string `json:"creator_id"`
	Pagination
}

type QueryItemRes struct {
	Data     []Collection `json:"data"`
	Metadata interface{}  `json:"metadata"`
}