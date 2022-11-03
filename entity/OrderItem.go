package entity

import "time"

type OrderItem struct {
	Id       uint64 `gorm:"primaryKey:autoIncreament;" json:"id"`
	Quantity uint   `json:"quantity"`
	Price    uint64

	ProductId uint64  `gorm:"not null" json:"-"`
	Product   Product `gorm:"foreign:ProductId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"product"`

	OrderId uint64 `gorm:"not null" json:"-"`
	Order   Order  `gorm:"foreign:OrderId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"order"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
