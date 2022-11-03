package entity

import "time"

type User struct {
	Id        	uint64 `gorm:"primaryKey:autoIncreament" json:"id"`
	Name      	string `gorm:"type:varchar(100);not null" json:"name"`
	Email     	string `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	Password  	string `gorm:"->;<-;not null" json:"-"`
	Token  		string `gorm:"-" json:"token,omitemtpy"`	

	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}