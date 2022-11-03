package router

import (
	"ecommerce-golang/controller"

	"github.com/gin-gonic/gin"
)

func NewProductRouter(r *gin.Engine, rGroup *gin.RouterGroup, c controller.ProductController) {
	// without middleware
	r.GET("products/", c.Index)
	r.GET("products/:id", c.Show)

	// with middleware
	rGroup.POST("products/", c.Store)
	rGroup.PUT("products/:id", c.Update)
	rGroup.DELETE("products/:id", c.Delete)
}
