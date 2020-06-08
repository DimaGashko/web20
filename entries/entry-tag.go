package entries

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Value string `gorm:"value"`
}
