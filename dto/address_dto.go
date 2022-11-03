package dto

type AddressInsertDTO struct {
	Province    string `json:"province" validate:"required" form:"province" bind:"required"`
	City        string `json:"city" validate:"required" form:"city" bind:"required"`
	Subdistrict string `json:"subdistrict" validate:"required" form:"subdistrict" bind:"required"`
	Address     string `json:"Address" validate:"required" form:"Address" bind:"required"`
	Zipcode     string `json:"zipcode" validate:"required" form:"zipcode" bind:"required"`
}

type AddressUpdateDTO struct {
	Id          uint64 `json:"id" validate:"required" form:"id" bind:"required"`
	Province    string `json:"province" validate:"required" form:"province" bind:"required"`
	City        string `json:"city" validate:"required" form:"city" bind:"required"`
	Subdistrict string `json:"subdistrict" validate:"required" form:"subdistrict" bind:"required"`
	Address     string `json:"Address" validate:"required" form:"Address" bind:"required"`
	Zipcode     string `json:"zipcode" validate:"required" form:"zipcode" bind:"required"`
}
