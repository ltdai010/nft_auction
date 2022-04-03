package repos

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"nft_auction/pkg/models"
)

func (r *RepoPG) LoginUser(ctx context.Context, user *models.Users) (err error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pubkey"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"last_login": gorm.Expr("CURRENT_TIMESTAMP")}),
	}).Create(user)
	return nil
}

func (r *RepoPG) GetUserProfile(ctx context.Context, pubkey string) (*models.Users, error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	ret := &models.Users{}
	if err := tx.Model(&models.Users{}).Where("pubkey = ?", pubkey).First(ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}
