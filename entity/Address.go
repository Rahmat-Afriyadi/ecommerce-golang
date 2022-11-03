package entity

import (
	"time"
)

type Address struct {
	Id          uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Province    string `gorm:"not null;type:varchar(50)" json:"province"`
	City        string `gorm:"not null;type:varchar(50)" json:"city"`
	Subdistrict string `gorm:"not null;type:varchar(50)" json:"subdistrict"`
	Address     string `gorm:"not null;type:text" json:"Address"`
	Zipcode     string `gorm:"not null;type:varchar(20)" json:"zipcode"`

	CustomerId uint64   `json:"-"`
	Customer   Customer `gorm:"foreign:CustomerId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
