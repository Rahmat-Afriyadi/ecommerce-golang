package entity

import "time"

type Customer struct {
	Id        uint64   `gorm:"primaryKey:autoIncreament" json:"id"`
	UserId    uint64   `gorm:"not null" json:"-"`
	User      User     `gorm:"foreign:UserId,constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Address   *Address `json:"address,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
