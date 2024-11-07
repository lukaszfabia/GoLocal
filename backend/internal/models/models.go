package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null:size:255"`
	LastName  string `gorm:"not null:size:255"`

	Email    string `gorm:"not null;size:100;unique"`
	Password string `gorm:"not null;size:400"`

	IsPremium bool    `gorm:"default:'false'"`
	ImageURL  *string `gorm:"null;size:1024"`
}
