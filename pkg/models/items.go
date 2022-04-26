package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Item struct {
	BaseModel
	ItemID       string          `json:"item_id" gorm:"index"`
	OwnerID      uuid.UUID       `json:"owner_id"`
	Owner        User            `json:"owner"`
	CreatorID    uuid.UUID       `json:"creator_id"`
	Creator      User            `json:"creator" gorm:"foreignKey:creator_id;references:id"`
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

type ItemReq struct {
	ItemID       string          `json:"item_id" gorm:"index"`
	OwnerID      uuid.UUID       `json:"owner_id"`
	CollectionID uuid.UUID       `json:"collection_id"`
	Metadata     json.RawMessage `json:"metadata" swaggertype:"object"`
	Category     string          `json:"category"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
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
	Name         string `json:"name" form:"name"`
	CollectionID string `json:"collection_id" form:"collection_id"`
	OwnerID      string `json:"owner_id" form:"owner_id"`
	LikedBy      string `json:"liked_by" form:"liked_by"`
	CreatorID    string `json:"creator_id" form:"creator_id"`
	Pagination
}

type QueryItemRes struct {
	Data     []Collection `json:"data"`
	Metadata interface{}  `json:"metadata"`
}
