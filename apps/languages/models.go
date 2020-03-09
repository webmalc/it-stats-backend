package languages

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

// GetLanguageSources returns the language sources
func GetLanguageSources() []string {
	return []string{"pypl", "tiobe"}
}

// Language is a programming language struct
type Language struct {
	gorm.Model
	Title    string `gorm:"size:100;not null;index:title"`
	Position uint   `gorm:"not null;index:position"`
	Source   string `gorm:"size:100;not null;index:source"`
}

// Validate validates struct
func (l *Language) Validate(db *gorm.DB) {
	if !govalidator.IsIn(l.Source, GetLanguageSources()...) {
		_ = db.AddError(
			validations.NewError(l, "Source", "the source is inccorrect"),
		)
	}
}
