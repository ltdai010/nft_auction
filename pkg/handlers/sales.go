package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"nft_auction/pkg/consts"
	"nft_auction/pkg/models"
	"nft_auction/pkg/services"
)

func NewSaleHandler(service services.SalesInterface) *SaleHandler {
	return &SaleHandler{service}
}

type SaleHandler struct {
	service services.SalesInterface
}

// Get
// @Tags Sales
// @Summary Get Sale
// @Description Get Sale
// @ID get-sale
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Sale
// @Router /sales/{id} [get]
func (c *SaleHandler) Get(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))

	user, err := c.service.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Query
// @Tags Sales
// @Summary Query Sale
// @Description Query Sale
// @ID query-sale
// @Accept  json
// @Produce  json
// @Param item_id query string false "creator_id"
// @Success 200 {object} models.QuerySaleRes
// @Router /sales/query/list [get]
func (c *SaleHandler) Query(ctx *gin.Context) {
	var req models.QuerySaleReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := c.service.Query(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Post
// @Tags Sales
// @Summary Post Sales
// @Description Post Sales
// @Security ApiKeyAuth
// @ID post-sales
// @Accept  json
// @Produce  json
// @Param data body models.SaleReq true "data"
// @Success 200 {object} models.Sale
// @Router /sales [post]
func (c *SaleHandler) Post(ctx *gin.Context) {
	var req models.SaleReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userId := uuid.MustParse(ctx.GetHeader("x-user-id"))
	user, err := c.service.Create(ctx, &req, &userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
