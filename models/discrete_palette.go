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

// DiscretePalette discrete palette
//
// swagger:model DiscretePalette
type DiscretePalette struct {

	// Array of colors in the palette
	Colors []string `json:"colors"`

	// Unique identity string
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Label for palette
	Label string `json:"label,omitempty"`

	// Type of palette
	Type string `json:"type,omitempty"`
}

// Validate validates this discrete palette
func (m *DiscretePalette) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this discrete palette based on the context it is used
func (m *DiscretePalette) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscretePalette) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscretePalette) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscretePalette) UnmarshalBinary(b []byte) error {
	var res DiscretePalette
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
