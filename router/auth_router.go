package router

import (
	"ecommerce-golang/controller"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(r *gin.Engine, c controller.AuthController) {
	r.POST("auth/sign-in", c.SignIn)
	r.POST("auth/sign-up", c.SignUp)
}
