package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"nft_auction/pkg/models"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

// Migrate
// @Tags Migration
// @Security ApiKeyAuth
// @Summary migrate
// @Description migrate
// @ID migrate
// @Accept  json
// @Produce  json
// @Success 200 {string} succes
// @Router /migrate [get]
func (h *MigrationHandler) Migrate(ctx *gin.Context) {

	_ = h.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := h.db.AutoMigrate([]interface{}{&models.User{}, &models.Collection{}, &models.Item{}, &models.ItemLike{}}...)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
