package entity

import "time"

type Product struct {
	Id          uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Title       string `gorm:"not null;type:varchar(100)" json:"title"`
	Description string `gorm:"not null;type:text" json:"description"`
	Image       string `gorm:"not null;type:varchar(100)" json:"image"`
	Price       uint64 `gorm:"not null;type:int(11)" json:"price"`
	Stock       uint64 `gorm:"not null;type:int(11)" json:"stock"`

	CategoryId uint64   `gorm:"not null" json:"-"`
	Category   Category `gorm:"foreign:CategoryId;constraint:onDelete:CASCADE,onUpdate:CASCADE" json:"Category"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
