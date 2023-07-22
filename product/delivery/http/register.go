package http

import (
	"github.com/dmytrodemianchuk/boxfit-backend/product"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc product.UseCase) {
	h := NewHandler(uc)

	products := router.Group("/products")
	{
		products.POST("", h.Create)
		products.GET("", h.Get)
		products.DELETE("", h.Delete)
	}
}
