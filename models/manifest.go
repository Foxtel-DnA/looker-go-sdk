// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Manifest manifest
//
// swagger:model Manifest
type Manifest struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Imports for a project
	// Read Only: true
	Imports []*ImportedProject `json:"imports"`

	// Localization settings for a project
	// Read Only: true
	LocalizationSettings *LocalizationSettings `json:"localization_settings,omitempty"`

	// Manifest project name
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this manifest
func (m *Manifest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizationSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manifest) validateImports(formats strfmt.Registry) error {
	if swag.IsZero(m.Imports) { // not required
		return nil
	}

	for i := 0; i < len(m.Imports); i++ {
		if swag.IsZero(m.Imports[i]) { // not required
			continue
		}

		if m.Imports[i] != nil {
			if err := m.Imports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Manifest) validateLocalizationSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalizationSettings) { // not required
		return nil
	}

	if m.LocalizationSettings != nil {
		if err := m.LocalizationSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localization_settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this manifest based on the context it is used
func (m *Manifest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalizationSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manifest) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Manifest) contextValidateImports(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "imports", "body", []*ImportedProject(m.Imports)); err != nil {
		return err
	}

	for i := 0; i < len(m.Imports); i++ {

		if m.Imports[i] != nil {
			if err := m.Imports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Manifest) contextValidateLocalizationSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalizationSettings != nil {
		if err := m.LocalizationSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localization_settings")
			}
			return err
		}
	}

	return nil
}

func (m *Manifest) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Manifest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Manifest) UnmarshalBinary(b []byte) error {
	var res Manifest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
