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

// CredentialsSaml credentials saml
//
// swagger:model CredentialsSaml
type CredentialsSaml struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Timestamp for the creation of this credential
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// EMail address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// Has this credential been disabled?
	// Read Only: true
	IsDisabled *bool `json:"is_disabled,omitempty"`

	// Timestamp for most recent login using credential
	// Read Only: true
	LoggedInAt string `json:"logged_in_at,omitempty"`

	// Saml IdP's Unique ID for this user
	// Read Only: true
	SamlUserID string `json:"saml_user_id,omitempty"`

	// Short name for the type of this kind of credential
	// Read Only: true
	Type string `json:"type,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this credentials saml
func (m *CredentialsSaml) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsSaml) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this credentials saml based on the context it is used
func (m *CredentialsSaml) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsDisabled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoggedInAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSamlUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsSaml) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CredentialsSaml) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", string(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "email", "body", string(m.Email)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateIsDisabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_disabled", "body", m.IsDisabled); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateLoggedInAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "logged_in_at", "body", string(m.LoggedInAt)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateSamlUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "saml_user_id", "body", string(m.SamlUserID)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsSaml) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsSaml) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsSaml) UnmarshalBinary(b []byte) error {
	var res CredentialsSaml
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
