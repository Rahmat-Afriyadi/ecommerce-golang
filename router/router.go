package router

import (
	"ecommerce-golang/app"
	"ecommerce-golang/config"
	"ecommerce-golang/controller"
	"ecommerce-golang/middleware"
	"ecommerce-golang/service"

	"github.com/gin-gonic/gin"
)

var (
	// middleware
	jwtService    service.JWTService = service.NewJWTServiceImpl(config.EnvSecretJwt())
	jwtMiddleware gin.HandlerFunc    = middleware.AuthorizeJWT(jwtService)

	// controller
	authController     controller.AuthController     = app.InitializedAuthController(config.EnvSecretJwt())
	productController  controller.ProductController  = app.InitializedProductController()
	categoryController controller.CategoryController = app.InitializedCategoryController()
)

func MainRouter(r *gin.Engine) {
	auth := r.Group("/", jwtMiddleware)
	NewAuthRouter(r, authController)
	NewProductRouter(r, auth, productController)
	NewCategoryRouter(r, auth, categoryController)
}
