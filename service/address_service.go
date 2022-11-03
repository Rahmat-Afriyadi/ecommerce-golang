package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"
	"ecommerce-golang/helper"
	"ecommerce-golang/repository"

	"github.com/go-playground/validator/v10"
)

type AddressServiceImpl struct {
	repository.AddressRepository
	*validator.Validate
}

func NewAddressServiceImpl(r repository.AddressRepository, v *validator.Validate) *AddressServiceImpl {
	return &AddressServiceImpl{
		r,
		v,
	}
}

func (service *AddressServiceImpl) FindById(id uint64) (entity.Address, error) {
	return service.AddressRepository.FindById(id) , nil
}
func (service *AddressServiceImpl) Create(dto dto.AddressInsertDTO) entity.Address {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	var Address entity.Address
	helper.FillStruct(dto, &Address)
	return service.AddressRepository.Insert(Address)
}
func (service *AddressServiceImpl) Update(dto dto.AddressUpdateDTO) entity.Address {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	var Address entity.Address
	helper.FillStruct(dto, &Address)
	return service.AddressRepository.Update(Address)
}
func (service *AddressServiceImpl) Delete(id uint64) {
	Address := service.AddressRepository.FindById(id)
	service.AddressRepository.Delete(Address)
}
