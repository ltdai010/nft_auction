package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft_auction/conf"
	"nft_auction/pkg/repos"
)

func NewRoute() *gin.Engine {
	route := gin.Default()
	route.Use(cors.Default())
	v1 := route.Group("/v1")
	users := v1.Group("/users/")
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", conf.LoadEnv().DBHost, conf.LoadEnv().DBUser, conf.LoadEnv().DBPass, conf.LoadEnv().DBName, conf.LoadEnv().DBPort)
	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Panicln(err)
	}

	repo := repos.NewPGRepo(db)

	// users
	InitUserRoutes(users, repo)

	// migration
	InitMigrationRoutes(v1, repo)

	return route
}
