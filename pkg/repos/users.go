package repos

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"nft_auction/pkg/models"
)

func (r *RepoPG) LoginUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pubkey"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"last_login": gorm.Expr("CURRENT_TIMESTAMP")}),
	}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RepoPG) GetUserProfile(ctx context.Context, id string) (*models.Users, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	ret := &models.Users{}
	if err := tx.Model(&models.Users{}).Where("id = ?", id).First(ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}
