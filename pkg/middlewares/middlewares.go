package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nft_auction/pkg/middlewares/jwttoken"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Request.Header.Set("pubkey", "")
		context.Request.Header.Set("x-user-id", "")
		tokenStr := context.GetHeader("authorization")
		if tokenStr == "" && context.Request.Method == http.MethodGet {
			return
		}
		if strings.Contains(context.Request.URL.String(), "login") {
			return
		}
		token := jwttoken.JwtTokenFrom(tokenStr)
		tokenInfo, err := token.ExtractTokenMetadata(&key.PublicKey)
		if err != nil || tokenInfo.Expired < time.Now().Unix() || tokenInfo.Pubkey == "" {
			log.Println("auth error", err)
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		context.Request.Header.Set("x-user-id", tokenInfo.UserID)
		context.Request.Header.Set("pubkey", tokenInfo.Pubkey)
	}
}
