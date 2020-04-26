package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(1);not null;unique"`
	Password  string `gorm:"type:size(255);not null"`
}
