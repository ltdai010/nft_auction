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

type Users struct {
	repo repos.PGInterface
}

type UserServiceInterface interface {
	Login(ctx context.Context, pubkey string) (*models.UsersLogin, error)
	RefreshToken(ctx context.Context, refreshToken string) (*models.LoginTokenResponse, error)
	GetProfile(ctx context.Context, id string) (*models.User, error)
}

func NewUserService(repo repos.PGInterface) UserServiceInterface {
	return &Users{
		repo: repo,
	}
}

func (s *Users) RefreshToken(ctx context.Context, refreshToken string) (*models.LoginTokenResponse, error) {
	token, err := middlewares.RefreshNewToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return &models.LoginTokenResponse{
		Token:        token.Token,
		ExpiredAt:    token.ExpiredAt,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Users) Login(ctx context.Context, pubkey string) (*models.UsersLogin, error) {
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

	user, err := s.repo.LoginUser(ctx, &models.User{
		Pubkey:  pubkey,
		Address: address.String(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token, err := middlewares.GenerateLoginToken(user.ID.String(), pubkey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.UsersLogin{
		Data:  *user,
		Token: *token,
	}, nil
}

func (s *Users) GetProfile(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.GetUserProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
