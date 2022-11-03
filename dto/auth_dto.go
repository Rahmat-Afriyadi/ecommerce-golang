package dto

type SignInDTO struct {
	Email string `json:"email" validate:"required,email" form:"email" bind:"required"`
	Password string `json:"password" validate:"required" form:"password" bind:"required"`
}

type SignUpDTO struct {
	Name string `json:"name" validate:"required" form:"name" bind:"name"`
	Email string `json:"email" validate:"required,email" form:"email" bind:"email"`
	Password string `json:"password" validate:"required" form:"password" bind:"password"`
}