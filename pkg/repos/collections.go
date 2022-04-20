package repos

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nft_auction/pkg/models"
)

func (r *RepoPG) CreateCollection(ctx context.Context, req *models.Collection) (*models.Collection, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Create(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return req, nil
}

func (r *RepoPG) GetCollection(ctx context.Context, cid *uuid.UUID) (*models.Collection, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	res := models.Collection{}
	if err := tx.Model(&models.Collection{}).Where("id = ?", cid).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RepoPG) UpdateCollection(ctx context.Context, req *models.Collection) (*models.Collection, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Model(&models.Collection{}).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return req, nil
}

func (r *RepoPG) DeleteCollection(ctx context.Context, cid *uuid.UUID) error {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	if err := tx.Delete(&models.Collection{BaseModel: models.BaseModel{ID: *cid}}).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *RepoPG) QueryCollections(ctx context.Context, req *models.QueryCollectionReq) (*models.QueryCollectionRes, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	tx = tx.Model(&models.Collection{})
	if req.CreatorID != "" {
		tx = tx.Where("creator_id = ?", req.CreatorID)
	}
	if req.Name != "" {
		tx = tx.Where("name ILIKE ?", "%"+req.Name+"%")
	}
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)
	collections := []models.Collection{}
	rs := models.QueryCollectionRes{}
	var total int64
	if err := tx.Count(&total).Limit(pageSize).Offset(r.GetOffset(page, pageSize)).Find(&collections).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	var err error
	if rs.Metadata, err = r.GetPaginationInfo("", nil, total, page, pageSize); err != nil {
		return nil, err
	}
	rs.Data = collections
	return &rs, nil
}
