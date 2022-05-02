package routes

import (
	"github.com/gin-gonic/gin"
	"nft_auction/pkg/handlers"
	"nft_auction/pkg/repos"
	"nft_auction/pkg/services"
)

func InitSalesRoutes(rg *gin.RouterGroup, repo repos.PGInterface) {
	service := services.NewSalesService(repo)
	handler := handlers.NewSaleHandler(service)
	rg.POST("", handler.Post)
	rg.GET("/:id", handler.Get)
	rg.GET("/query/list", handler.Query)
	rg.PUT("/actions/buy", handler.Buy)
}
