package routes

import (
	"github.com/gin-gonic/gin"
	"nft_auction/pkg/handlers"
	"nft_auction/pkg/repos"
)

func InitMigrationRoutes(rg *gin.RouterGroup, repo repos.PGInterface) {
	handler := handlers.NewMigrationHandler(repo.DB())
	rg.GET("/migrate", handler.Migrate)
}
