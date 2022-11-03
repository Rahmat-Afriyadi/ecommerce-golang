package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"
	"ecommerce-golang/helper"
	"ecommerce-golang/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	repository.CategoryRepository
	*validator.Validate
}

func NewCategoryServiceImpl(r repository.CategoryRepository, v *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		r,
		v,
	}
}

var _ CategoryService = (*CategoryServiceImpl)(nil)

func (service *CategoryServiceImpl) All() []entity.Category {
	return service.All()
}
func (service *CategoryServiceImpl) FindById(id uint64) entity.Category {
	return service.FindById(id)
}
func (service *CategoryServiceImpl) Create(dto dto.CategoryInsertDTO) (entity.Category, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		return entity.Category{}, err
	}
	var Category entity.Category
	helper.FillStruct(dto, &Category)
	if dto.Image != nil {
		imageUrl, err := helper.ImageUploadHelper(dto.Image)
		if err != nil {
			return entity.Category{}, err
		}
		Category.Image = imageUrl
	}	
	return service.CategoryRepository.Insert(Category), nil
}
func (service *CategoryServiceImpl) Update(dto dto.CategoryUpdateDTO) (entity.Category, error) {
	err := service.Validate.Struct(dto)
	if err != nil {
		return entity.Category{}, err
	}

	var Category entity.Category
	helper.FillStruct(dto, &Category)
	return service.CategoryRepository.Update(Category), nil
}
func (service *CategoryServiceImpl) Delete(id uint64) {
	Category := service.CategoryRepository.FindById(id)
	service.CategoryRepository.Delete(Category)
}
