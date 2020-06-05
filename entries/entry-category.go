package entries

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Slug string `gorm:"slug;unique_index"`
	Name string `gorm:"name"`
}
