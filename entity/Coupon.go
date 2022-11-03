package entity

import "time"

type Coupon struct {
	Id        uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Percent   uint8  `gorm:"not null;type:int(3)" json:"percent"`
	Max       uint64 `gorm:"not null;type:int(11)" json:"max"`
	Image     string `gorm:"type:varchar(100)" json:"image"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
