package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	PostID  uint
	User    User `gorm:"foreignKey:UserID"`
	Post    Post `gorm:"foreignKey:PostID"`
}
