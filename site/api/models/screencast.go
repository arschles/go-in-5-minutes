package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Screencast struct {
	ID              uuid.UUID `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Title           string    `json:"title" db:"title"`
	Date            time.Time `json:"date" db:"date"`
	Summary         string    `json:"summary" db:"summary"`
	Complete        string    `json:"complete" db:"complete"`
	YouTubeEmbedURL string    `json:"youtube_embed_url" db:"youtube_embed_url"`
}

func (s Screencast) HumanTime() string {
	return s.Date.Format("Mon Jan 2, 2006")
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Screencast) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Title, Name: "Title"},
		&validators.StringIsPresent{Field: s.Summary, Name: "Summary"},
		&validators.StringIsPresent{Field: s.Complete, Name: "Complete"},
		&validators.StringIsPresent{Field: s.YouTubeEmbedURL, Name: "YouTubeEmbedURL"},
		// &validators.TimeIsPresent{Field: s.Date, Name:"Date"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Screencast) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Screencast) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
