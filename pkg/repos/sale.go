package repos

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nft_auction/pkg/models"
)

func (r *RepoPG) CreateSale(ctx context.Context, req *models.Sale) (*models.Sale, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Save(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (r *RepoPG) GetSale(ctx context.Context, id *uuid.UUID) (*models.Sale, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	var req models.Sale
	if err := tx.Model(&models.Sale{}).Where("id = ?", id).First(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return &req, nil
}

func (r *RepoPG) QuerySales(ctx context.Context, req *models.QuerySaleReq) (*models.QuerySaleRes, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	tx = tx.Model(&models.Sale{})
	var data []models.Sale
	var res models.QuerySaleRes
	if err := tx.Model(&models.Sale{}).Where("item_id = ?", req.ItemID).Find(&data).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	res.Data = data
	return &res, nil
}

func (r *RepoPG) UpdateSale(ctx context.Context, req *models.Sale) (*models.Sale, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Save(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return req, nil
}

func (r *RepoPG) DeleteSale(ctx context.Context, id *uuid.UUID) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Model(&models.Sale{}).Where("id = ?", id).Delete(&models.Sale{}).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
