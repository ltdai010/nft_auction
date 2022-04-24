package services

import (
	"context"
	"github.com/google/uuid"
	"nft_auction/pkg/models"
	"nft_auction/pkg/repos"
)

type Items struct {
	repo repos.PGInterface
}

func NewItemsService(repo repos.PGInterface) ItemsInterface {
	return &Items{repo: repo}
}

type ItemsInterface interface {
	Create(ctx context.Context, item *models.Item) (*models.Item, error)
	Get(ctx context.Context, cid *uuid.UUID) (*models.Item, error)
	Update(ctx context.Context, item *models.Item) (*models.Item, error)
	Like(ctx context.Context, itemId *uuid.UUID, uid *uuid.UUID) error
	Delete(ctx context.Context, cid *uuid.UUID) error
	Query(ctx context.Context, req *models.QueryItemReq) (*models.QueryItemRes, error)
}

func (c *Items) Create(ctx context.Context, Item *models.Item) (*models.Item, error) {
	return c.repo.CreateItem(ctx, Item)
}

func (c *Items) Get(ctx context.Context, cid *uuid.UUID) (*models.Item, error) {
	return c.repo.GetItem(ctx, cid)
}

func (c *Items) Update(ctx context.Context, Item *models.Item) (*models.Item, error) {
	return c.repo.UpdateItem(ctx, Item)
}

func (c *Items) Delete(ctx context.Context, cid *uuid.UUID) error {
	return c.repo.DeleteItem(ctx, cid)
}

func (c *Items) Query(ctx context.Context, req *models.QueryItemReq) (*models.QueryItemRes, error) {
	return c.repo.QueryItems(ctx, req)
}

func (c *Items) Like(ctx context.Context, itemId *uuid.UUID, uid *uuid.UUID) error {
	return c.repo.Like(ctx, &models.ItemLike{
		BaseModel: models.BaseModel{
			CreatorID: *uid,
		},
		ItemID: *itemId,
		UserID: *uid,
	})
}
