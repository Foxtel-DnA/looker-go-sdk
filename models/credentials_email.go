// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CredentialsEmail credentials email
//
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
	ForcedPasswordResetAtNextLogin bool `json:"forced_password_reset_at_next_login,omitempty"`

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

// ContextValidate validate this credentials email based on the context it is used
func (m *CredentialsEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsDisabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoggedInAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePasswordResetURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsEmail) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CredentialsEmail) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", string(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidateIsDisabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_disabled", "body", m.IsDisabled); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidateLoggedInAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "logged_in_at", "body", string(m.LoggedInAt)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidatePasswordResetURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "password_reset_url", "body", string(m.PasswordResetURL)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsEmail) contextValidateUserURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_url", "body", strfmt.URI(m.UserURL)); err != nil {
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
