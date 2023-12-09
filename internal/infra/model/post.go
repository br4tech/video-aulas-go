package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title   string `gorm:column:title`
	Content string `gorm:column:content`

	UserID int
	User   User `gorm:foreignKey:UserID`
}
