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

// ModelSet model set
//
// swagger:model ModelSet
type ModelSet struct {

	// all access
	// Read Only: true
	AllAccess *bool `json:"all_access,omitempty"`

	// built in
	// Read Only: true
	BuiltIn *bool `json:"built_in,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// models
	Models []string `json:"models"`

	// Name of ModelSet
	Name string `json:"name,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this model set
func (m *ModelSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelSet) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this model set based on the context it is used
func (m *ModelSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBuiltIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *ModelSet) contextValidateAllAccess(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "all_access", "body", m.AllAccess); err != nil {
		return err
	}

	return nil
}

func (m *ModelSet) contextValidateBuiltIn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "built_in", "body", m.BuiltIn); err != nil {
		return err
	}

	return nil
}

func (m *ModelSet) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelSet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ModelSet) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSet) UnmarshalBinary(b []byte) error {
	var res ModelSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
