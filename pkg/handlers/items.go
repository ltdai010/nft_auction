package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"nft_auction/pkg/consts"
	"nft_auction/pkg/models"
	"nft_auction/pkg/services"
)

func NewItemHandler(service services.ItemsInterface) *ItemHandler {
	return &ItemHandler{service}
}

type ItemHandler struct {
	service services.ItemsInterface
}

// Get
// @Tags Items
// @Summary Get Item
// @Description Get Item
// @ID get-item
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (c *ItemHandler) Get(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))

	user, err := c.service.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Query
// @Tags Items
// @Summary Get Item
// @Description Get Item
// @ID query-items
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Param collection_id query string false "collection_id"
// @Param liked_by query string false "liked by"
// @Param owner_id query string false "owner_id"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Success 200 {object} models.QueryItemRes
// @Router /items/query/list [get]
func (c *ItemHandler) Query(ctx *gin.Context) {
	var req models.QueryItemReq
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
// @Tags Items
// @Summary Post Item
// @Description Post Item
// @Security ApiKeyAuth
// @ID post-item
// @Accept  json
// @Produce  json
// @Param data body models.Item true "data"
// @Success 200 {object} models.Item
// @Router /items [post]
func (c *ItemHandler) Post(ctx *gin.Context) {
	var req models.ItemReq
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

// Like
// @Tags Items
// @Summary Like Item
// @Description Like Item
// @Security ApiKeyAuth
// @ID put-item
// @Accept  json
// @Produce  json
// @Param data body models.Item true "data"
// @Success 200 {string} success
// @Router /items/action/like/{id} [put]
func (c *ItemHandler) Like(ctx *gin.Context) {
	itemId := uuid.MustParse(ctx.Param("id"))
	userId := uuid.MustParse(ctx.GetHeader("x-user-id"))
	err := c.service.Like(ctx, &itemId, &userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// Put
// @Tags Items
// @Summary Put Item
// @Description Put Item
// @Security ApiKeyAuth
// @ID put-item
// @Accept  json
// @Produce  json
// @Param data body models.Item true "data"
// @Success 200 {object} models.Item
// @Router /items/{id} [put]
func (c *ItemHandler) Put(ctx *gin.Context) {
	var req models.ItemReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id := uuid.MustParse(ctx.Param("id"))
	userId := uuid.MustParse(ctx.GetHeader("x-user-id"))
	user, err := c.service.Update(ctx, &userId, &req, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Delete
// @Tags Items
// @Summary Delete Item
// @Description Delete Item
// @Security ApiKeyAuth
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} success
// @Router /items/{id} [delete]
func (c *ItemHandler) Delete(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))
	err := c.service.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
