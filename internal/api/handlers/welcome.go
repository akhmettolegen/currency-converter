package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) Welcome(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"message": "success"})
}
