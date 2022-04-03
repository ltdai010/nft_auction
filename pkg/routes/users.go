package routes

import (
	"github.com/gin-gonic/gin"
	"nft_auction/pkg/handlers"
	"nft_auction/pkg/repos"
	"nft_auction/pkg/services"
)

func InitUserRoutes(rg *gin.RouterGroup, repo repos.PGInterface) {
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	rg.POST("/login", handler.Login)
	rg.GET("/:id", handler.GetProfile)
}
