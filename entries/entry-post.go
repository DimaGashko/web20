package entries

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Slug        string `gorm:"slug;unique_index;not null" json:"slug"`
	Title       string `gorm:"title;not null" json:"title"`
	Image       string `gorm:"image;not null" json:"image"`
	Description string `gorm:"description;not null" json:"description"`
	Content     Md     `gorm:"content;not null" json:"content"`
	Author      string `gorm:"author;" json:"author"`
	Listed      bool   `gorm:"listed;not null" json:"listed"`
	Secret      string `gorm:"secret;not null" json:"secret"`
	Category    string `gorm:"category;not null" json:"category"`
}
