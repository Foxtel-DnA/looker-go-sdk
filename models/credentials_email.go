// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CredentialsEmail credentials email
// swagger:model CredentialsEmail
type CredentialsEmail struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Timestamp for the creation of this credential
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// EMail address used for user login
	Email string `json:"email,omitempty"`

	// Force the user to change their password upon their next login
	// Read Only: true
	ForcedPasswordResetAtNextLogin *bool `json:"forced_password_reset_at_next_login,omitempty"`

	// Has this credential been disabled?
	// Read Only: true
	IsDisabled *bool `json:"is_disabled,omitempty"`

	// Timestamp for most recent login using credential
	// Read Only: true
	LoggedInAt string `json:"logged_in_at,omitempty"`

	// Url with one-time use secret token that the user can use to reset password
	// Read Only: true
	PasswordResetURL string `json:"password_reset_url,omitempty"`

	// Short name for the type of this kind of credential
	// Read Only: true
	Type string `json:"type,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Link to get this user
	// Read Only: true
	// Format: uri
	UserURL strfmt.URI `json:"user_url,omitempty"`
}

// Validate validates this credentials email
func (m *CredentialsEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsEmail) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) validateUserURL(formats strfmt.Registry) error {

	if swag.IsZero(m.UserURL) { // not required
		return nil
	}

	if err := validate.FormatOf("user_url", "body", "uri", m.UserURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsEmail) UnmarshalBinary(b []byte) error {
	var res CredentialsEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
