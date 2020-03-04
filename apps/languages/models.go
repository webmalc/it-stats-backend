package languages

import (
	"github.com/jinzhu/gorm"
)

// Language is a programming language struct
type Language struct {
	gorm.Model
	Title    string `gorm:"size:100;not null;index:title"`
	Position uint   `gorm:"not null;index:position"`
	Source   string `gorm:"size:100;not null;index:source"`
}
