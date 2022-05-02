package services

import (
	"context"
	"github.com/google/uuid"
	"nft_auction/pkg/models"
	"nft_auction/pkg/repos"
)

type Sales struct {
	repo repos.PGInterface
}

func NewSalesService(repo repos.PGInterface) SalesInterface {
	return &Sales{repo: repo}
}

type SalesInterface interface {
	Create(ctx context.Context, sale *models.SaleReq, creatorID *uuid.UUID) (*models.Sale, error)
	Buy(ctx context.Context, saleId *uuid.UUID, creatorID *uuid.UUID) (*models.Sale, error)
	Get(ctx context.Context, cid *uuid.UUID) (*models.Sale, error)
	Query(ctx context.Context, req *models.QuerySaleReq) (*models.QuerySaleRes, error)
	Delete(ctx context.Context, cid *uuid.UUID) error
}

func (s *Sales) Create(ctx context.Context, sale *models.SaleReq, creatorID *uuid.UUID) (*models.Sale, error) {
	// update status item
	item, err := s.repo.GetItem(ctx, &sale.ItemID)
	if err != nil {
		return nil, err
	}
	item.Status = models.OnSale
	item, err = s.repo.UpdateItem(ctx, item)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateSale(ctx, &models.Sale{
		BaseModel: models.BaseModel{
			CreatorID: *creatorID,
		},
		ItemID:         sale.ItemID,
		Price:          sale.Price,
		CoinBuy:        sale.CoinBuy,
		CoinBuyAddress: sale.CoinBuyAddress,
		Decimal:        sale.Decimal,
	})
}

func (s *Sales) Query(ctx context.Context, req *models.QuerySaleReq) (*models.QuerySaleRes, error) {
	return s.repo.QuerySales(ctx, req)
}

func (s *Sales) Get(ctx context.Context, cid *uuid.UUID) (*models.Sale, error) {
	return s.repo.GetSale(ctx, cid)
}

func (s *Sales) Delete(ctx context.Context, cid *uuid.UUID) error {
	return s.repo.DeleteSale(ctx, cid)
}

func (s *Sales) Buy(ctx context.Context, saleId *uuid.UUID, creatorID *uuid.UUID) (*models.Sale, error) {
	sale, err := s.repo.GetSale(ctx, saleId)
	if err != nil {
		return nil, err
	}
	sale.BuyerID = creatorID
	sale, err = s.repo.UpdateSale(ctx, sale)
	if err != nil {
		return nil, err
	}
	item, err := s.repo.GetItem(ctx, &sale.ItemID)
	if err != nil {
		return nil, err
	}
	item.OwnerID = *creatorID
	item.Status = models.Unset
	item, err = s.repo.UpdateItem(ctx, item)
	if err != nil {
		return nil, err
	}
	return sale, nil
}
