package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// adminCabinet godoc
// @Summary Admin personal cabinet for managing the site
// @Tags admin
// @Security ApiKeyAuth
// @Router /api/v1/admin [get]
func (h *Handler) adminCabinetV1(c *gin.Context) {
	fmt.Println("adminCabinetV1")
}
