package handlers

import (
	"github.com/akhmettolegen/currency-converter/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Convert(ctx *gin.Context) {

	var body model.RequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, err.Error())
	}

	response := h.App.Manager.Converter(&body)

	ctx.JSON(200, response)
}
