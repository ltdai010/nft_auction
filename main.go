package main

import (
	"log"
	"nft_auction/conf"
	"nft_auction/pkg/routes"
	"os"
)

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
