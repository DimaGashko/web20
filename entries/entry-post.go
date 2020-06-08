package entries

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Slug        string   `gorm:"slug;unique_index" json:"slug"`
	Title       string   `gorm:"title" json:"title"`
	Image       string   `gorm:"image" json:"image"`
	Description string   `gorm:"description" json:"description"`
	Content     Md       `gorm:"content" json:"content"`
	Author      string   `gorm:"author" json:"author"`
	Listed      bool     `gorm:"listed" json:"listed"`
	Secret      string   `gorm:"secret" json:"secret"`
	Category    Category `gorm:"category;foreignkey:Category" json:"category"`
	Tags        []Tag
}
