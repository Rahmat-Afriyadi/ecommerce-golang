package entity

import "time"

type CommentRate struct {
	Id         uint64   `gorm:"primaryKey:autoIncreatment" json:"id"`
	Rating     uint     `json:"rating"`
	Comment    string   `gorm:"text" json:"comment"`
	ProductId  uint64   `gorm:"not null" json:"-"`
	Product    Product  `gorm:"foreign:ProductId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"product"`
	CustomerId uint64   `gorm:"not null" josn:"-"`
	Customer   Customer `gorm:"foreign:CustomerId;constraint:onUpdate:CASCADE,onDelete:CASCADE" josn:"customer"`
	Image      string   `json:"image"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
