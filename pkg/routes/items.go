package routes

import (
	"github.com/gin-gonic/gin"
	"nft_auction/pkg/handlers"
	"nft_auction/pkg/repos"
	"nft_auction/pkg/services"
)

func InitItemsRoutes(rg *gin.RouterGroup, repo repos.PGInterface) {
	service := services.NewItemsService(repo)
	handler := handlers.NewItemHandler(service)
	rg.POST("/", handler.Post)
	rg.PUT("/items/action/like/:id", handler.Like)
	rg.GET("/:id", handler.Get)
	rg.GET("/query/list", handler.Query)
	rg.PUT("/:id", handler.Put)
	rg.DELETE("/:id", handler.Delete)
}
