package services

import (
	"context"
	"github.com/google/uuid"
	"nft_auction/pkg/models"
	"nft_auction/pkg/repos"
)

type Collections struct {
	repo repos.PGInterface
}

func NewCollectionsService(repo repos.PGInterface) CollectionsInterface {
	return &Collections{repo: repo}
}

type CollectionsInterface interface {
	Create(ctx context.Context, collection *models.CollectionReq, creatorId *uuid.UUID) (*models.Collection, error)
	Get(ctx context.Context, cid *uuid.UUID) (*models.Collection, error)
	Update(ctx context.Context, id *uuid.UUID, collection *models.CollectionReq, updaterId *uuid.UUID) (*models.Collection, error)
	Delete(ctx context.Context, cid *uuid.UUID) error
	Query(ctx context.Context, req *models.QueryCollectionReq) (*models.QueryCollectionRes, error)
}

func (c *Collections) Create(ctx context.Context, collection *models.CollectionReq, creatorId *uuid.UUID) (*models.Collection, error) {
	return c.repo.CreateCollection(ctx, &models.Collection{
		BaseModel: models.BaseModel{
			CreatorID: *creatorId,
		},
		Name:     collection.Name,
		Address:  collection.Address,
		Metadata: collection.Metadata,
	})
}

func (c *Collections) Get(ctx context.Context, cid *uuid.UUID) (*models.Collection, error) {
	return c.repo.GetCollection(ctx, cid)
}

func (c *Collections) Update(ctx context.Context, id *uuid.UUID, collection *models.CollectionReq, updater *uuid.UUID) (*models.Collection, error) {
	collectionData, err := c.repo.GetCollection(ctx, id)
	if err != nil {
		return nil, err
	}
	collectionData.UpdaterID = *updater
	collectionData.Name = collection.Name
	collectionData.Address = collection.Address
	collectionData.Metadata = collection.Metadata
	return c.repo.UpdateCollection(ctx, collectionData)
}

func (c *Collections) Delete(ctx context.Context, cid *uuid.UUID) error {
	return c.repo.DeleteCollection(ctx, cid)
}

func (c *Collections) Query(ctx context.Context, req *models.QueryCollectionReq) (*models.QueryCollectionRes, error) {
	return c.repo.QueryCollections(ctx, req)
}
