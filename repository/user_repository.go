package repository

import (
	"ecommerce-golang/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Conn: db,
	}
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

func (r *UserRepositoryImpl) All() []entity.User {
	var Users []entity.User
	r.Conn.Find(&Users)
	return Users
}

func (r *UserRepositoryImpl) FindById(id uint64) entity.User {
	var User entity.User
	r.Conn.Find(&User, id)
	return User
}
func (r *UserRepositoryImpl) FindByEmail(email string) entity.User {
	var u = entity.User{
		Email: email,
	}
	r.Conn.Find(&u)
	return u
}
func (r *UserRepositoryImpl) Insert(e entity.User) entity.User {
	e.Password = r.hashAndSalt([]byte(e.Password))
	r.Conn.Save(&e)
	return e
	
}

func (r *UserRepositoryImpl) Update(e entity.User) entity.User {
	r.Conn.Save(&e)
	return e
}
func (r *UserRepositoryImpl) Delete(e entity.User) {
	r.Conn.Delete(&e)
}

func (r *UserRepositoryImpl) hashAndSalt(passByte []byte) string {
	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}
