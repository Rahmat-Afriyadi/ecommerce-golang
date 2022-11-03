package entity

import "time"

type Order struct {
	Id          uint            `gorm:"primaryKey:autoIncreament" json:"id"`
	CustomerId  uint64          `gorm:"not null" json:"-"`
	Customer    Customer        `gorm:"foreign:CustomerId;constraint:onDelete:CASCADE,onUpdate:CASCADE" json:"customer"`
	CouponId    uint64          `json:"-"`
	Coupon      Coupon          `gorm:"foreign:CouponId;constraint:onDelete:SET NULL,onUpdate:CASCADE" json:"coupon,omitempty"`
	ShippingId  uint64          `json:"-"`
	Shipping    ShippingAddress `gorm:"foreign:ShippingId;constraint:onDelete:SET NULL,onUpdate:CASCADE" json:"shipping_address,omitempty"`
	OrderItems  []OrderItem     `json:"order_items,omitEmpty"`
	ReceiptCode string          `json:"receipt_code"`
	Complete    bool            `json:"complete"`
	TotalPrice  uint            `gorm:"<-:update" json:"total_price"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
