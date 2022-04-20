package middlewares

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"nft_auction/conf"
	"nft_auction/pkg/middlewares/jwttoken"
	"nft_auction/pkg/models"
	"os"
	"time"
)

var key *rsa.PrivateKey

func init() {
	k, err := GetKeyFromFile()
	if err != nil {
		log.Panicln(err)
	}
	key = k
}

func PubkeyFromSign(signature string) (pubkey string, err error) {
	hash := accounts.TextHash([]byte(conf.Message))
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return "", err
	}
	sig[crypto.RecoveryIDOffset] -= 27
	sigPublicKey, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sigPublicKey), nil
}

func GetKeyFromFile() (*rsa.PrivateKey, error) {
	f, err := os.OpenFile("./rsa.private", os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func RefreshNewToken(refreshToken string) (*models.NewTokenResponse, error) {
	rToken := jwttoken.JwtRefreshTokenFrom(refreshToken)
	rTokenInfo, err := rToken.ExtractTokenMetadata(&key.PublicKey)
	if err != nil {
		return nil, err
	}
	claim := jwt.MapClaims{
		"userId": rTokenInfo.UserID,
		"pubkey": rTokenInfo.Pubkey,
		"exp":    time.Now().UTC().Add(time.Minute * 60).Unix(),
	}
	tokenInfo := jwttoken.TokenInfoFromClaims(claim)
	token, err := tokenInfo.Create(key)
	if err != nil {
		return nil, err
	}
	return &models.NewTokenResponse{
		Token:     token,
		ExpiredAt: tokenInfo.Expired,
	}, nil
}

func GenerateLoginToken(userId, pubkey string) (*models.LoginTokenResponse, error) {
	rClaim := jwt.MapClaims{
		"userId": userId,
		"pubkey": pubkey,
	}
	claim := jwt.MapClaims{
		"userId": userId,
		"pubkey": pubkey,
		"exp":    time.Now().UTC().Add(time.Minute * 60).Unix(),
	}
	rTokenInfo := jwttoken.RefreshTokenInfoFromClaims(rClaim)
	rToken, err := rTokenInfo.Create(key)
	if err != nil {
		return nil, err
	}
	tokenInfo := jwttoken.TokenInfoFromClaims(claim)
	token, err := tokenInfo.Create(key)
	if err != nil {
		return nil, err
	}

	return &models.LoginTokenResponse{
		Token:        token,
		RefreshToken: rToken,
		ExpiredAt:    tokenInfo.Expired,
	}, nil
}
