package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft_auction/pkg/consts"
	"nft_auction/pkg/middlewares"
	"nft_auction/pkg/models"
	"nft_auction/pkg/services"
)

func NewUserHandler(userService services.UserServiceInterface) *UserController {
	return &UserController{userService}
}

type UserController struct {
	userService services.UserServiceInterface
}

func (c *UserController) Login(ctx *gin.Context) {
	req := &models.UserLoginRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, consts.ErrBadRequest)
		return
	}
	pubkey, err := middlewares.PubkeyFromSign(req.Signature)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, consts.ErrBadRequest)
		return
	}
	user, err := c.userService.Login(ctx, pubkey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, consts.ErrInternalServer)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetProfile(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
