package repository

import (
	"ecommerce-golang/entity"

	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	Conn *gorm.DB
}

func NewAddressRepositoryImpl(c *gorm.DB) *AddressRepositoryImpl {
	return &AddressRepositoryImpl{
		c,
	}
}

var _ AddressRepository = (*AddressRepositoryImpl)(nil)

func (r *AddressRepositoryImpl) FindById(id uint64) entity.Address {
	Address := entity.Address{}
	r.Conn.Find(&Address, id)
	return Address
}
func (r *AddressRepositoryImpl) Insert(e entity.Address) entity.Address {
	r.Conn.Save(&e)
	return e
}
func (r *AddressRepositoryImpl) Update(e entity.Address) entity.Address {
	r.Conn.Save(&e)
	return e
}
func (r *AddressRepositoryImpl) Delete(e entity.Address) {
	r.Conn.Delete(&e)
}
