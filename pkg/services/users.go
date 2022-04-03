package services

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"nft_auction/pkg/middlewares"
	"nft_auction/pkg/models"
	"nft_auction/pkg/repos"
)

type UserService struct {
	repo repos.PGInterface
}

type UserServiceInterface interface {
	Login(ctx context.Context, pubkey string) (*models.UsersLogin, error)
	GetProfile(ctx context.Context, pubkey string) (*models.Users, error)
}

func NewUserService(repo repos.PGInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, pubkey string) (*models.UsersLogin, error) {
	pb, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pk, err := crypto.UnmarshalPubkey(pb)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	address := crypto.PubkeyToAddress(*pk)

	err = s.repo.LoginUser(ctx, &models.Users{
		Pubkey:  pubkey,
		Address: address.String(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user, err := s.repo.GetUserProfile(ctx, pubkey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token, err := middlewares.GenerateLoginToken(pubkey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.UsersLogin{
		Data:  *user,
		Token: *token,
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, pubkey string) (*models.Users, error) {
	user, err := s.repo.GetUserProfile(ctx, pubkey)
	if err != nil {
		return nil, err
	}
	return user, nil
}
