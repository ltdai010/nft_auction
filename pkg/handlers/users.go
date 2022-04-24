package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nft_auction/pkg/consts"
	"nft_auction/pkg/middlewares"
	"nft_auction/pkg/models"
	"nft_auction/pkg/services"
)

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{userService}
}

type UserHandler struct {
	userService services.UserServiceInterface
}

// Login
// @Tags User
// @Summary Login user
// @Description Login user
// @ID login
// @Accept  json
// @Produce  json
// @Param data body models.UserLoginRequest true "body"
// @Success 200 {string} success
// @Router /users/login [post]
func (c *UserHandler) Login(ctx *gin.Context) {
	req := &models.UserLoginRequest{}
	if err := ctx.BindJSON(req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, consts.ErrBadRequest)
		return
	}
	pubkey, err := middlewares.PubkeyFromSign(req.Signature)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, consts.ErrBadRequest)
		return
	}
	user, err := c.userService.Login(ctx, pubkey)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, consts.ErrInternalServer)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetProfile
// @Tags User
// @Summary Get profile
// @Description Get profile
// @ID get-profile
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.User
// @Router /users/profile/{id} [get]
func (c *UserHandler) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetProfile(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
