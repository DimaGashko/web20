package entries

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Slug        string   `gorm:"slug;unique_index"`
	Title       string   `gorm:"title"`
	Image       string   `gorm:"image"`
	Description string   `gorm:"description"`
	Content     Md       `gorm:"content"`
	Author      string   `gorm:"author"`
	Listed      bool     `gorm:"listed"`
	Category    Category `gorm:"category;foreignkey:Category"`
	Tags        []Tag
}
