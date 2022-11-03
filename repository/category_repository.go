package repository

import (
	"ecommerce-golang/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	Conn *gorm.DB
}

func NewCategoryRepositoryImpl(c *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		c,
	}
}

var _ CategoryRepository = (*CategoryRepositoryImpl)(nil)

func (r *CategoryRepositoryImpl) FindById(id uint64) entity.Category {
	Category := entity.Category{}
	r.Conn.Find(&Category, id)
	return Category
}
func (r *CategoryRepositoryImpl) Insert(e entity.Category) entity.Category {
	r.Conn.Save(&e)
	return e
}
func (r *CategoryRepositoryImpl) Update(e entity.Category) entity.Category {
	r.Conn.Save(&e)
	return e
}
func (r *CategoryRepositoryImpl) Delete(e entity.Category) {
	r.Conn.Delete(&e)
}
