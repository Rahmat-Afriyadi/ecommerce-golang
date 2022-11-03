package repository

import (
	"ecommerce-golang/entity"

	"gorm.io/gorm"
)

type CouponRepositoryImpl struct {
	Conn *gorm.DB
}

func NewCouponRepositoryImpl(c *gorm.DB) *CouponRepositoryImpl {
	return &CouponRepositoryImpl{
		c,
	}
}

func (r *CouponRepositoryImpl) FindById(id uint64) entity.Coupon {
	Coupon := entity.Coupon{}
	r.Conn.Find(&Coupon, id)
	return Coupon
}
func (r *CouponRepositoryImpl) Insert(e entity.Coupon) entity.Coupon {
	r.Conn.Save(&e)
	return e
}
func (r *CouponRepositoryImpl) Update(e entity.Coupon) entity.Coupon {
	r.Conn.Save(&e)
	return e
}
func (r *CouponRepositoryImpl) Delete(e entity.Coupon) {
	r.Conn.Delete(&e)
}
