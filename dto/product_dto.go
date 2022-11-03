package dto

import "mime/multipart"

type ProductInsertDTO struct {
	Title       string         `json:"title" validate:"required" form:"title" bind:"required"`
	Description string         `json:"description" validate:"required" form:"description" bind:"required"`
	Image       multipart.File `json:"image" validate:"required" form:"image" bind:"required"`
	Price       uint64         `json:"price" validate:"required" form:"price" bind:"required"`
	Stock       uint64         `json:"stock" validate:"required" form:"stock" bind:"required"`
	CategoryId  uint64         `json:"category_id" validate:"required" form:"category_id" bind:"required"`
}

type ProductUpdateDTO struct {
	Id          uint64 `json:"id" validate:"required" form:"id" bind:"required"`
	Title       string `json:"title" validate:"required" form:"title" bind:"required"`
	Description string `json:"description" validate:"required" form:"description" bind:"required"`
	Image       []byte `json:"image"`
	Price       uint64 `json:"price" validate:"required" form:"price" bind:"required"`
	Stock       uint64 `json:"stock" validate:"required" form:"stock" bind:"required"`
	CategoryId  uint64 `json:"category_id" validate:"required" form:"category_id" bind:"required"`
}
