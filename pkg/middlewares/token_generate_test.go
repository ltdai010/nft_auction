package middlewares

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"nft_auction/conf"
	"testing"
)

func TestPubkeyFromSign(t *testing.T) {
	hash := accounts.TextHash([]byte(conf.Message))
	privk, err := crypto.HexToECDSA("32b8395b72d438bf64da8d8eb6f2eb9dfda0d6ea07e18773614813d0c63ca95b")
	if err != nil {
		t.Error(err)
		return
	}
	bb := crypto.FromECDSAPub(&privk.PublicKey)
	t.Log(hex.EncodeToString(bb))
	t.Log(hex.EncodeToString(hash))
	sig, _ := crypto.Sign(hash, privk)
	t.Log(hex.EncodeToString(sig))
	pk, err := PubkeyFromSign("f279004777bb7501ac01820f66da22bf192e373c52ee1771f8843d71775fa0691ec5fa2f207b1be0452da1a94905ee6a55077608369373295ae97dcc58976fff1b")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(pk)
	b, err := hex.DecodeString(pk)
	if err != nil {
		t.Fatal(err)
	}
	pb, err := crypto.UnmarshalPubkey(b)
	if err != nil {
		t.Fatal(err)
	}
	address := crypto.PubkeyToAddress(*pb).String()
	t.Log(address)
}
