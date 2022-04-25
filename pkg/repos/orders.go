package repos

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nft_auction/pkg/models"
)

func (r *RepoPG) CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Create(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (r *RepoPG) GetOrder(ctx context.Context, id *uuid.UUID) (*models.Order, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	var res models.Order
	if err := tx.Model(&models.Order{}).Where("id = ?", id.String()).First(&res).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func (r *RepoPG) QueryOrders(ctx context.Context, req *models.QueryOrderReq) (*models.QueryOrderRes, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	tx = tx.Model(&models.Order{})
	if req.ItemID != "" {
		tx = tx.Where("item_id = ?", req.ItemID)
	}
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	var total int64
	var res models.QueryOrderRes
	if err := tx.Model(&models.Order{}).Count(&total).Limit(pageSize).Offset(r.GetOffset(page, pageSize)).First(&res).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	res.Metadata = r.GetPaginationInfo(total, page, pageSize)
	return &res, nil
}

func (r *RepoPG) UpdateOrder(ctx context.Context, req *models.Order) (*models.Order, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Save(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (r *RepoPG) DeleteOrder(ctx context.Context, cid *uuid.UUID) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Model(&models.Order{}).Where("id = ?", cid).Delete(&models.Order{}).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
