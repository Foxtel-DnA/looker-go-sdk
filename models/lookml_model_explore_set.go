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

// LookmlModelExploreSet lookml model explore set
//
// swagger:model LookmlModelExploreSet
type LookmlModelExploreSet struct {

	// Name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Value set
	// Read Only: true
	Value []string `json:"value"`
}

// Validate validates this lookml model explore set
func (m *LookmlModelExploreSet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this lookml model explore set based on the context it is used
func (m *LookmlModelExploreSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModelExploreSet) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreSet) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", []string(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreSet) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
