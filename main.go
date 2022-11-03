package main

import (
	"ecommerce-golang/app"
	"ecommerce-golang/config"
	"ecommerce-golang/controller"
	"ecommerce-golang/middleware"
	"ecommerce-golang/router"
	"ecommerce-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	db *gorm.DB           = config.SetupDatabaseConnection()
	v  validator.Validate = *validator.New()

	// middleware
	jwtService    service.JWTService = service.NewJWTServiceImpl(config.EnvSecretJwt())
	jwtMiddleware gin.HandlerFunc    = middleware.AuthorizeJWT(jwtService)

	// controller
	authController     controller.AuthController     = app.InitializedAuthController(config.EnvSecretJwt())
	productController  controller.ProductController  = app.InitializedProductController()
	categoryController controller.CategoryController = app.InitializedCategoryController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.StaticFS("/static", http.Dir("resources"))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "test")
	})

	router.MainRouter(r)

	r.Run(":8080")
}
