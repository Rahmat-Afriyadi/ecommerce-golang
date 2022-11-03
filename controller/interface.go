package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type AuthController interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}

type ProductController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
