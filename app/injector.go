//go:build wireinject
// +build wireinject

package app

import (
	"ecommerce-golang/config"
	"ecommerce-golang/controller"
	"ecommerce-golang/repository"
	"ecommerce-golang/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var productSet = wire.NewSet(
	repository.NewProductRepositoryImpl,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
	service.NewProductServiceImpl,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
	controller.NewProductControllerImpl,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

func InitializedProductController() controller.ProductController {
	wire.Build(
		config.SetupDatabaseConnection,
		validator.New,
		productSet,
	)
	return nil
}

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedCategoryController() controller.CategoryController {
	wire.Build(
		config.SetupDatabaseConnection,
		validator.New,
		categorySet,
	)
	return nil
}

var authSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewJWTServiceImpl,
	wire.Bind(new(service.JWTService), new(*service.JwtServiceImpl)),
	service.NewAuthServiceImpl,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	controller.NewAuthControllerImpl,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

func InitializedAuthController(secret string) controller.AuthController {
	wire.Build(
		config.SetupDatabaseConnection,
		validator.New,
		authSet,
	)
	return nil
}
