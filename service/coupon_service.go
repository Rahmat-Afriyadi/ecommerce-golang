package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"
	"ecommerce-golang/helper"
	"ecommerce-golang/repository"

	"github.com/go-playground/validator/v10"
)

type CouponServiceImpl struct {
	repository.CouponRepository
	*validator.Validate
}

func NewCouponServiceImpl(r repository.CouponRepository, v *validator.Validate) *CouponServiceImpl {
	return &CouponServiceImpl{
		r,
		v,
	}
}

func (service *CouponServiceImpl) FindById(id uint64) entity.Coupon {
	return service.FindById(id)
}
func (service *CouponServiceImpl) Create(dto dto.CouponInsertDTO) entity.Coupon {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	var Coupon entity.Coupon
	helper.FillStruct(dto, &Coupon)
	return service.CouponRepository.Insert(Coupon)
}
func (service *CouponServiceImpl) Update(dto dto.CouponUpdateDTO) entity.Coupon {
	err := service.Validate.Struct(dto)
	if err != nil {
		panic(err.Error())
	}
	var Coupon entity.Coupon
	helper.FillStruct(dto, &Coupon)
	return service.CouponRepository.Update(Coupon)
}
func (service *CouponServiceImpl) Delete(id uint64) {
	Coupon := service.CouponRepository.FindById(id)
	service.CouponRepository.Delete(Coupon)
}
