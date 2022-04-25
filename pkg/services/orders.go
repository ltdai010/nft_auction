package services

import (
	"context"
	"github.com/google/uuid"
	"nft_auction/pkg/models"
	"nft_auction/pkg/repos"
)

type Orders struct {
	repo repos.PGInterface
}

func NewOrdersService(repo repos.PGInterface) OrdersInterface {
	return &Orders{repo: repo}
}

type OrdersInterface interface {
	Create(ctx context.Context, order *models.Order) (*models.Order, error)
	Get(ctx context.Context, cid *uuid.UUID) (*models.Order, error)
	Update(ctx context.Context, order *models.Order) (*models.Order, error)
	Delete(ctx context.Context, cid *uuid.UUID) error
	Query(ctx context.Context, req *models.QueryOrderReq) (*models.QueryOrderRes, error)
}

func (o Orders) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o Orders) Get(ctx context.Context, cid *uuid.UUID) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o Orders) Update(ctx context.Context, order *models.Order) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o Orders) Delete(ctx context.Context, cid *uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (o Orders) Query(ctx context.Context, req *models.QueryOrderReq) (*models.QueryOrderRes, error) {
	//TODO implement me
	panic("implement me")
}
