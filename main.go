package main

import (
	"log"
	"nft_auction/conf"
	_ "nft_auction/docs"
	"nft_auction/pkg/routes"
	"os"
)

// @title Finan Loyalty API
// @version 1.0
// @description This is Finan Loyalty server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @Host localhost:8080
// @BasePath  /v1
func main() {
	conf.SetEnv()

	//app := routes.NewService()
	//	//ctx := context.Background()
	//	//err := app.Start(ctx)
	//	//if err != nil {
	//	//	logger.Tag("main").Error(err)
	//	//}
	err := routes.NewRoute().Run("0.0.0.0:8080")
	if err != nil {
		log.Println(err)
	}
	os.Clearenv()
}
