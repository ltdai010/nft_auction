package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"nft_auction/pkg/consts"
	"nft_auction/pkg/models"
	"nft_auction/pkg/services"
)

func NewCollectionHandler(service services.CollectionsInterface) *CollectionHandler {
	return &CollectionHandler{service}
}

type CollectionHandler struct {
	service services.CollectionsInterface
}

// Get
// @Tags Collections
// @Summary Get Collection
// @Description Get Collection
// @ID get-collection
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Collection
// @Router /collections/{id} [get]
func (c *CollectionHandler) Get(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))

	user, err := c.service.Get(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Query
// @Tags Collections
// @Summary Get Collection
// @Description Get Collection
// @ID query-collection
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Param creator_id query string false "creator_id"
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Success 200 {object} models.QueryCollectionRes
// @Router /collections/query/list [get]
func (c *CollectionHandler) Query(ctx *gin.Context) {
	var req models.QueryCollectionReq
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
// @Tags Collections
// @Summary Post Collection
// @Description Post Collection
// @Security ApiKeyAuth
// @ID post-collection
// @Accept  json
// @Produce  json
// @Param data body models.CollectionReq true "data"
// @Success 200 {object} models.Collection
// @Router /collections [post]
func (c *CollectionHandler) Post(ctx *gin.Context) {
	var req models.CollectionReq
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

// Put
// @Tags Collections
// @Summary Put Collection
// @Description Put Collection
// @Security ApiKeyAuth
// @ID put-collection
// @Accept  json
// @Produce  json
// @Param data body models.CollectionReq true "data"
// @Success 200 {object} models.Collection
// @Router /collections/{id} [put]
func (c *CollectionHandler) Put(ctx *gin.Context) {
	var req models.CollectionReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userId := uuid.MustParse(ctx.GetHeader("x-user-id"))
	id := uuid.MustParse(ctx.Param("id"))
	user, err := c.service.Update(ctx, &id, &req, &userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Delete
// @Tags Collections
// @Summary Delete Collection
// @Description Delete Collection
// @Security ApiKeyAuth
// @ID delete-collection
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} success
// @Router /collections/{id} [delete]
func (c *CollectionHandler) Delete(ctx *gin.Context) {
	id := uuid.MustParse(ctx.Param("id"))
	err := c.service.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, consts.ErrNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
