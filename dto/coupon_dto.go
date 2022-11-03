package dto

type CouponInsertDTO struct {
	Percent uint8  `json:"percent" validate:"required,max=100" form:"percent" bind:"required"`
	Max     uint64 `json:"max" validate:"required" form:"max" bind:"required"`
	Image   string `json:"image"`
}

type CouponUpdateDTO struct {
	Percent uint8  `json:"percent" validate:"required,max=100" form:"percent" bind:"required"`
	Max     uint64 `json:"max" validate:"required" form:"max" bind:"required"`
	Image   string `json:"image"`
}
