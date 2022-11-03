package repository

import (
	"ecommerce-golang/entity"
)

type AddressRepository interface {
	FindById(id uint64) entity.Address
	Insert(e entity.Address) entity.Address
	Update(e entity.Address) entity.Address
	Delete(e entity.Address)
}

type CouponRepository interface {
	FindById(id uint64) entity.Coupon
	Insert(e entity.Coupon) entity.Coupon
	Update(e entity.Coupon) entity.Coupon
	Delete(e entity.Coupon)
}

type CategoryRepository interface {
	FindById(id uint64) entity.Category
	Insert(e entity.Category) entity.Category
	Update(e entity.Category) entity.Category
	Delete(e entity.Category)
}

type ProductRepository interface {
	All() []entity.Product
	FindById(id uint64) entity.Product
	Insert(e entity.Product) entity.Product
	Update(e entity.Product) entity.Product
	Delete(e entity.Product) error
	Where(*entity.Product) entity.Product
}

type UserRepository interface {
	All() []entity.User
	FindById(id uint64) entity.User
	FindByEmail(email string) entity.User
	Insert(e entity.User) entity.User
	Update(e entity.User) entity.User
	Delete(e entity.User)
	hashAndSalt(passByte []byte) string
}
