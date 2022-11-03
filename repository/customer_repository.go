package repository

import (
	"ecommerce-golang/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Insert(userId uint64) entity.Customer
}

type CustomerRepositoryImpl struct {
	Conn *gorm.DB
}

func NewCustomerRepositoryImpl(c *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		c,
	}
}

var _ ProductRepository = (*ProductRepositoryImpl)(nil)

func (r *CustomerRepositoryImpl) Insert(userId uint64) entity.Customer {
	Customer := entity.Customer{UserId: userId}
	r.Conn.Save(&Customer)
	return Customer
}
