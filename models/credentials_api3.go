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

// CredentialsApi3 credentials api3
//
// swagger:model CredentialsApi3
type CredentialsApi3 struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// API key client_id
	// Read Only: true
	ClientID string `json:"client_id,omitempty"`

	// Timestamp for the creation of this credential
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Has this credential been disabled?
	// Read Only: true
	IsDisabled *bool `json:"is_disabled,omitempty"`

	// Short name for the type of this kind of credential
	// Read Only: true
	Type string `json:"type,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this credentials api3
func (m *CredentialsApi3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsApi3) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this credentials api3 based on the context it is used
func (m *CredentialsApi3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsDisabled(ctx, formats); err != nil {
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

func (m *CredentialsApi3) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CredentialsApi3) contextValidateClientID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "client_id", "body", string(m.ClientID)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsApi3) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", string(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsApi3) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsApi3) contextValidateIsDisabled(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_disabled", "body", m.IsDisabled); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsApi3) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *CredentialsApi3) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsApi3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsApi3) UnmarshalBinary(b []byte) error {
	var res CredentialsApi3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
