package dto

import "mime/multipart"

type CategoryInsertDTO struct {
	Title       string `json:"title" validate:"required" form:"title" bind:"required"`
	Description string `json:"description" validate:"required" form:"description" bind:"required"`
	Image       multipart.File `json:"image" `
}

type CategoryUpdateDTO struct {
	Id          uint64 `json:"id" validate:"required" form:"id" bind:"required"`
	Title       string `json:"title" validate:"required" form:"title" bind:"required"`
	Description string `json:"description" validate:"required" form:"description" bind:"required"`
	Image       multipart.File `json:"image" `
}
