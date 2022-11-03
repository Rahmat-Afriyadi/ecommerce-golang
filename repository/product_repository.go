package repository

import (
	"ecommerce-golang/entity"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Conn *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		Conn: db,
	}
}

var _ ProductRepository = (*ProductRepositoryImpl)(nil)

func (r *ProductRepositoryImpl) All() []entity.Product {
	var products []entity.Product
	r.Conn.Find(&products)
	return products
}
func (r *ProductRepositoryImpl) FindById(id uint64) entity.Product {
	var product entity.Product
	r.Conn.Find(&product, id)
	return product
}
func (r *ProductRepositoryImpl) Insert(e entity.Product) entity.Product {
	r.Conn.Save(&e)
	return e
}

func (r *ProductRepositoryImpl) Update(e entity.Product) entity.Product {
	r.Conn.Save(&e)
	return e
}
func (r *ProductRepositoryImpl) Delete(e entity.Product) error {
	r.Conn.Delete(&e)
	return nil
}
func (r *ProductRepositoryImpl) Where(e *entity.Product) entity.Product {
	var product entity.Product
	r.Conn.Where(&e).First(&product)
	return product
}
