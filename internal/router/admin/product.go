package admin

import (
	"goflow/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *handler.ProductHandler) {
	rg.GET("/products", h.List)
	rg.GET("/products/:id", h.Get)
	rg.POST("/products", h.Create)
	rg.PUT("/products/:id", h.Update)
	rg.DELETE("/products/:id", h.Delete)
}
