package controller

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/helper"
	"ecommerce-golang/service"
	"net/http"
	// "os/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	service.CategoryService
}

func NewCategoryControllerImpl(s service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		s,
	}
}

func (controller *CategoryControllerImpl) Index(ctx *gin.Context) {
	helper.BuildResponse(true, "all data Categorys", controller.CategoryService.All())
	return
}
func (controller *CategoryControllerImpl) Show(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	helper.BuildResponse(true, "detail Category by id", controller.CategoryService.FindById(id))
}
func (controller *CategoryControllerImpl) Store(ctx *gin.Context) {
	dto := dto.CategoryInsertDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	file, _, _ := ctx.Request.FormFile("image")
	defer file.Close()
	if file != nil {
		dto.Image = file
	}
	result, err := controller.CategoryService.Create(dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusAccepted, response)
	return
}

func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {
	dto := dto.CategoryUpdateDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		panic(err.Error())
	}
	result, err := controller.CategoryService.Update(dto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed To Process Request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	controller.CategoryService.Delete(id)
	res := helper.BuildResponse(true, "success", "successfully deleted Category")
	ctx.JSON(http.StatusAccepted, res)
	return
}
