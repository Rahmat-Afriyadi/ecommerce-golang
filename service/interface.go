package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"

	"github.com/dgrijalva/jwt-go"
)

type AddressService interface {
	Create(dto dto.AddressInsertDTO) entity.Address
	Update(dto dto.AddressUpdateDTO) entity.Address
	Delete(id uint64)
}

type AuthService interface {
	SignUp(SignUpDTO dto.SignUpDTO) (entity.User, error)
	VerifyCredential(SignInDTO dto.SignInDTO) (entity.User, error)
}

type CategoryService interface {
	All() []entity.Category
	FindById(id uint64) entity.Category
	Create(dto dto.CategoryInsertDTO) (entity.Category, error)
	Update(dto dto.CategoryUpdateDTO) (entity.Category, error)
	Delete(id uint64)
}

type CouponService interface {
	FindById(id uint64) entity.Coupon
	Create(dto dto.CouponInsertDTO) entity.Coupon
	Update(dto dto.CouponUpdateDTO) entity.Coupon
	Delete(id uint64)
}

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type ProductService interface {
	All() ([]entity.Product, error)
	FindById(id uint64) (entity.Product, error)
	Create(dto dto.ProductInsertDTO) (entity.Product, error)
	Update(dto dto.ProductUpdateDTO) (entity.Product, error)
	Delete(id uint64) error
}
