package controller

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/helper"
	"ecommerce-golang/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	service.ProductService
}

func NewProductControllerImpl(s service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		s,
	}
}

func (controller *ProductControllerImpl) Index(ctx *gin.Context) {
	res, err := controller.ProductService.All()
	if err != nil {
		helper.BuildErrorResponse("Failed to process request", err.Error(), nil)
	}
	helper.BuildResponse(true, "all data products", res)
	return
}

func (controller *ProductControllerImpl) Show(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	res, err := controller.ProductService.FindById(id)
	if err != nil {
		helper.BuildErrorResponse("Failed to process request", err.Error(), nil)
	}
	helper.BuildResponse(true, "detail product by id", res)
}

func (controller *ProductControllerImpl) Store(ctx *gin.Context) {
	dto := dto.ProductInsertDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	_, err = ctx.FormFile("image")
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	result, err := controller.ProductService.Create(dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusAccepted, response)
	return

}
func (controller *ProductControllerImpl) Update(ctx *gin.Context) {
	dto := dto.ProductUpdateDTO{}
	err := ctx.ShouldBind(&dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	file, _, _ := ctx.Request.FormFile("image")
	if file != nil {
		buff := make([]byte, 512)
		file.Read(buff)
		dto.Image = buff
	}
	result, err := controller.ProductService.Update(dto)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusAccepted, response)
	return
}
func (controller *ProductControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	err := controller.ProductService.Delete(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "success", "successfully deleted product")
	ctx.JSON(http.StatusAccepted, res)
	return
}
