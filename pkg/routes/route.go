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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRoute() *gin.Engine {
	cors.Default()
	route := gin.Default()
	route.Use(CORSMiddleware())
	v1 := route.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	users := v1.Group("/users/")
	collections := v1.Group("/collections/")
	items := v1.Group("/items/")
	sales := v1.Group("/sales/")
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
	InitSalesRoutes(sales, repo)

	// migration
	InitMigrationRoutes(v1, repo)

	return route
}
