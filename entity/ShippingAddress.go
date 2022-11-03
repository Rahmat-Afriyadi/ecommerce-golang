package entity

import "time"

type ShippingAddress struct {
	Id uint64 `gorm:"primaryKey:autoIncreament" json:"id"`

	CustomerId uint64   `json:"-"`
	Customer   Customer `gorm:"foreign:CustomerId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"Customer"`
	AddressId  uint64   `json:"-"`
	Address    Address  `gorm:"foreign:AddressId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"Address"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
