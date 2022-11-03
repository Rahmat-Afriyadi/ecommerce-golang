package service

import (
	"ecommerce-golang/dto"
	"ecommerce-golang/entity"
	"ecommerce-golang/helper"
	"ecommerce-golang/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthServiceImpl(ur repository.UserRepository, v *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{
		ur,
		v,
	}
}

var _ AuthService = (*AuthServiceImpl)(nil)

func (service *AuthServiceImpl) VerifyCredential(signInDTO dto.SignInDTO) (entity.User, error) {
	err := service.Validate.Struct(signInDTO)
	if err != nil {
		panic(err.Error())
	}
	user := service.UserRepository.FindByEmail(signInDTO.Email)
	if user.Id == 0 {
		return entity.User{}, errors.New("Email not found")
	}
	if valid := comparedPassword(user.Password, signInDTO.Password); valid == false {
		return entity.User{}, errors.New("worng Password")
	}
	return user, nil
}

func (service *AuthServiceImpl) SignUp(signUpDTO dto.SignUpDTO) (entity.User, error) {
	err := service.Validate.Struct(signUpDTO)
	if err != nil {
		return entity.User{}, err
	}
	if user := service.UserRepository.FindByEmail(signUpDTO.Email); user.Id != 0 {
		return entity.User{}, errors.New("email already exists")
	}
	user := entity.User{}
	helper.FillStruct(signUpDTO, &user)
	user.Password = signUpDTO.Password
	return service.UserRepository.Insert(user), nil
}

func comparedPassword(hashed string, plainPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainPass))
	if err != nil {
		return false
	}
	return true
}
