package api

import (
	"github.com/akhmettolegen/currency-converter/internal/api/handlers"
	"github.com/akhmettolegen/currency-converter/pkg/application"
	"github.com/gin-gonic/gin"
)

func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()

	handler := handlers.Get(&app)

	v1 := router.Group("v1")
	welcome := v1.Group("")
	{
		welcome.GET("/welcome", handler.Welcome)
	}

	return router, nil
}
