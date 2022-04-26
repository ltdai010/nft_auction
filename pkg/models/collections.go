package models

import "encoding/json"

type Collection struct {
	BaseModel
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Metadata *json.RawMessage `json:"metadata" swaggertype:"object"`
	Items    []Item           `json:"items"`
}

func (user *Collection) TableName() string {
	return "collections"
}

type CollectionReq struct {
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Metadata *json.RawMessage `json:"metadata" swaggertype:"object"`
}

type QueryCollectionReq struct {
	Name      string `json:"name" form:"name"`
	CreatorID string `json:"creator_id" form:"creator_id"`
	Pagination
}

type QueryCollectionRes struct {
	Data     []Collection `json:"data"`
	Metadata interface{}  `json:"metadata"`
}
