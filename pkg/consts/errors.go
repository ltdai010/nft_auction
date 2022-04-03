package consts

import "nft_auction/pkg/models"

var ErrBadRequest = models.NewErr(400, "Bad Request")
var ErrNotFound = models.NewErr(404, "Not found")
var ErrInternalServer = models.NewErr(500, "Internal Server Error")
