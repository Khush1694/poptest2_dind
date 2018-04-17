package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Bio       string    `json:"bio" db:"bio"`
	Location  string    `json:"location" db:"location"`
	Image     string    `json:"image" db:"image"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Title, Name: "Title"},
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName"},
		&validators.StringIsPresent{Field: u.Bio, Name: "Bio"},
		&validators.StringIsPresent{Field: u.Location, Name: "Location"},
		&LocationValidator{Field: u.Location, Name: "Location"},
		&validators.URLIsPresent{Field: u.Image, Name: "Image"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// The validation struct to hold the field name, field value, and any messages
type LocationValidator struct {
	Name    string
	Field   string
	Message string
}

// IsValid performs the validation based on city, state format
func (v *LocationValidator) IsValid(errors *validate.Errors) {
	parts := strings.Split(v.Field, ",")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		if v.Message == "" {
			v.Message = fmt.Sprintf("%s is not a valid location.", v.Name)
		}
		errors.Add(validators.GenerateKey(v.Name), v.Message)
	} else if len(parts) == 2 {
		state := strings.TrimSpace(parts[1])
		// Check that domain is valid
		if len(state) < 2 {
			if v.Message == "" {
				v.Message = fmt.Sprintf("%s does not provide a valid state or state abbreviation.", v.Name)
			}
			errors.Add(validators.GenerateKey(v.Name), v.Message)
		}
	}
}

func (u *User) AfterSave(tx *pop.Connection) error {
	tokened := strings.Split(u.Image, "/")
	fileName := tokened[len(tokened)-1]

	output, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer output.Close()

	response, err := http.Get(u.Image)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}
	return nil
}
