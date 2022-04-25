package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft_auction/conf"
	"nft_auction/pkg/middlewares"
	"nft_auction/pkg/repos"
)

func NewRoute() *gin.Engine {
	route := gin.Default()
	route.Use(cors.Default())
	v1 := route.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	v1.Use(cors.Default())
	users := v1.Group("/users/")
	collections := v1.Group("/collections/")
	items := v1.Group("/items/")
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", conf.LoadEnv().DBHost, conf.LoadEnv().DBUser, conf.LoadEnv().DBPass, conf.LoadEnv().DBName, conf.LoadEnv().DBPort)
	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Panicln(err)
	}

	repo := repos.NewPGRepo(db)

	v1.GET("swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// routes
	InitUserRoutes(users, repo)
	InitCollectionsRoutes(collections, repo)
	InitItemsRoutes(items, repo)

	// migration
	InitMigrationRoutes(v1, repo)

	return route
}
