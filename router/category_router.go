package router

import (
	"ecommerce-golang/controller"

	"github.com/gin-gonic/gin"
)

func NewCategoryRouter(r *gin.Engine, rGroup *gin.RouterGroup, c controller.CategoryController) {
	// without middleware
	r.GET("categories/", c.Index)
	r.GET("categories/:id", c.Show)

	// with middleware
	rGroup.POST("categories/", c.Store)
	rGroup.PUT("categories/:id", c.Update)
	rGroup.DELETE("categories/:id", c.Delete)
}
