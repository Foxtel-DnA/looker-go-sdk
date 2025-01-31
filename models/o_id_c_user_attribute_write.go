// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OIDCUserAttributeWrite o ID c user attribute write
//
// swagger:model OIDCUserAttributeWrite
type OIDCUserAttributeWrite struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Name of User Attribute in OIDC
	Name string `json:"name,omitempty"`

	// Required to be in OIDC assertion for login to be allowed to succeed
	Required bool `json:"required,omitempty"`

	// Looker User Attribute Ids
	UserAttributeIds []int64 `json:"user_attribute_ids"`
}

// Validate validates this o ID c user attribute write
func (m *OIDCUserAttributeWrite) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this o ID c user attribute write based on the context it is used
func (m *OIDCUserAttributeWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OIDCUserAttributeWrite) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *OIDCUserAttributeWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCUserAttributeWrite) UnmarshalBinary(b []byte) error {
	var res OIDCUserAttributeWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
