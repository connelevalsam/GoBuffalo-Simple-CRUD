package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Lecturer struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	School     string    `json:"school" db:"school"`
	FullName   string    `json:"full_name" db:"full_name"`
	Email      string    `json:"email" db:"email"`
	Department string    `json:"department" db:"department"`
	Age        int       `json:"age" db:"age"`
}

// String is not required by pop and may be deleted
func (l Lecturer) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Lecturers is not required by pop and may be deleted
type Lecturers []Lecturer

// String is not required by pop and may be deleted
func (l Lecturers) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (l *Lecturer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.School, Name: "School"},
		&validators.StringIsPresent{Field: l.FullName, Name: "FullName"},
		&validators.StringIsPresent{Field: l.Email, Name: "Email"},
		&validators.StringIsPresent{Field: l.Department, Name: "Department"},
		&validators.IntIsPresent{Field: l.Age, Name: "Age"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (l *Lecturer) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (l *Lecturer) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
