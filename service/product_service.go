package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"
	"ecommerce-golang/helper"
	"ecommerce-golang/repository"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	repository.ProductRepository
	*validator.Validate
}

func NewProductServiceImpl(r repository.ProductRepository, v *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		r,
		v,
	}
}

var _ ProductService = (*ProductServiceImpl)(nil)

func (service *ProductServiceImpl) All() ([]entity.Product, error) {
	return service.ProductRepository.All(), nil
}

func (service *ProductServiceImpl) FindById(id uint64) (entity.Product, error) {

	return service.FindById(id)
}

func (service *ProductServiceImpl) Create(dto dto.ProductInsertDTO) (entity.Product, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	uploadUrl, err := helper.ImageUploadHelper(dto.Image)
	if err != nil {
		return entity.Product{}, err
	}
	product := entity.Product{}
	helper.FillStruct(dto, &product)
	product.Image = uploadUrl
	return service.ProductRepository.Insert(product), nil
}

func (service *ProductServiceImpl) Update(dto dto.ProductUpdateDTO) (entity.Product, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	product := entity.Product{}
	helper.FillStruct(dto, &product)
	if dto.Image != nil {
		if helper.CheckImageType(http.DetectContentType(dto.Image)) == false {
			return entity.Product{}, errors.New("type of image unsuported")
		}
	}
	return service.ProductRepository.Update(product), nil
}

func (service *ProductServiceImpl) Delete(id uint64) error {
	product := service.ProductRepository.FindById(id)
	service.ProductRepository.Delete(product)
	return nil
}
