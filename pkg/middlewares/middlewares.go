package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nft_auction/pkg/middlewares/jwttoken"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("pubkey", "")
		tokenStr := context.GetHeader("authorization")
		if tokenStr == "" && context.Request.Method == http.MethodGet {
			return
		}
		token := jwttoken.JwtTokenFrom(tokenStr)
		tokenInfo, err := token.ExtractTokenMetadata(&key.PublicKey)
		if err != nil || tokenInfo.Expired < time.Now().Unix() || tokenInfo.Pubkey == "" {
			log.Println("auth error", err.Error())
			context.Abort()
			return
		}
		context.Header("pubkey", tokenInfo.Pubkey)
	}
}
