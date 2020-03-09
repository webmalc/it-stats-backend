package languages

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-backend/common/db"
)

// Should return the language sources
func TestGetLanguageSources(t *testing.T) {
	s := GetLanguageSources()
	assert.Equal(t, []string{"pypl", "tiobe"}, s)
}

// Should validate
func TestLanguage_Validate(t *testing.T) {
	conn := db.NewConnection()
	l := &Language{
		Title:    "go",
		Position: 1,
		Source:   "invalid",
	}
	l.Validate(conn.DB)
	errs := conn.GetErrors()

	assert.Error(t, errs[0])
	assert.Contains(t, errs[0].Error(), "source")
	conn.Close()

	conn = db.NewConnection()
	defer conn.Close()
	l.Source = "pypl"
	l.Validate(conn.DB)
	errs = conn.GetErrors()

	assert.Empty(t, errs)
}
