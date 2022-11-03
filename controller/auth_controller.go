package controller

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/helper"
	"ecommerce-golang/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthControllerImpl(s service.AuthService, jwt service.JWTService) *AuthControllerImpl {
	return &AuthControllerImpl{
		s,
		jwt,
	}
}

var _ AuthController = (*AuthControllerImpl)(nil)

func (controller *AuthControllerImpl) SignIn(ctx *gin.Context) {
	signInDto := dto.SignInDTO{}
	err := ctx.ShouldBind(&signInDto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res, err := controller.AuthService.VerifyCredential(signInDto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res.Token = controller.JWTService.GenerateToken(strconv.FormatUint(res.Id, 10))
	response := helper.BuildResponse(true, "success to login", res)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	return

}
func (controller *AuthControllerImpl) SignUp(ctx *gin.Context) {
	signUpDTO := dto.SignUpDTO{}
	err := ctx.ShouldBind(&signUpDTO)	
	if err != nil {
		res := helper.BuildErrorResponse("failed to proccess request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	user, err := controller.AuthService.SignUp(signUpDTO)
	if err != nil {
		res := helper.BuildErrorResponse("failed to proccess request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return	
	}
	res := helper.BuildResponse(true, "successfully register", user)
	ctx.JSON(http.StatusAccepted, res)
	return
}


