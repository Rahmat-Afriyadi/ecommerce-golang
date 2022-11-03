package entity

import "time"

type Category struct {
	Id          uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Title       string `gorm:"not null;type:varchar(100)" json:"title"`
	Description string `gorm:"not null;type:text" json:"description"`
	Image 		string `json:"image"`

	Products []Product `json:"products,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
