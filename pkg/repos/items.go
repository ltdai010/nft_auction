package repos

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nft_auction/pkg/models"
)

func (r *RepoPG) CreateItem(ctx context.Context, req *models.Item) (*models.Item, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Create(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return req, nil
}

func (r *RepoPG) GetItem(ctx context.Context, id *uuid.UUID) (*models.Item, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	res := models.Item{}
	if err := tx.Model(&models.Item{}).Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RepoPG) QueryItems(ctx context.Context, req *models.QueryItemReq) (*models.QueryItemRes, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	tx = tx.Model(&models.Item{})
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	items := []models.Item{}
	rs := models.QueryItemRes{}
	var total int64
	if req.CollectionID != "" {
		tx = tx.Where("collection_id = ?", req.CollectionID)
	}
	if req.OwnerID != "" {
		tx = tx.Where("owner_id = ?", req.OwnerID)
	}
	if req.Name != "" {
		tx = tx.Where("name = ?", "%"+req.Name+"%")
	}
	if err := tx.Count(&total).Limit(pageSize).Offset(r.GetOffset(page, pageSize)).Find(&items).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	var err error
	if rs.Metadata, err = r.GetPaginationInfo("", nil, total, page, pageSize); err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *RepoPG) UpdateItem(ctx context.Context, req *models.Item) (*models.Item, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Save(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return req, nil
}

func (r *RepoPG) DeleteItem(ctx context.Context, cid *uuid.UUID) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Delete(&models.Collection{BaseModel: models.BaseModel{ID: *cid}}).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
