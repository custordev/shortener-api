package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string 
	Role string `gorm:"default:USER;not null"`
	Link []Link
}

type Link struct{
	gorm.Model
	ShortCode string `gorm:"unique"`
	OriginalUrl string
	Clicks uint
	UserID uint 
	Favicon string

}