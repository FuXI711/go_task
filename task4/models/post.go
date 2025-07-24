package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint `gorm:"foreignKey:UserID"`
	User     User
	Comments []Comment
}
