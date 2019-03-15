// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIVersion Api version
// swagger:model ApiVersion
type APIVersion struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Current version for this Looker instance
	// Read Only: true
	CurrentVersion *APIVersionElement `json:"current_version,omitempty"`

	// Current Looker release version number
	// Read Only: true
	LookerReleaseVersion string `json:"looker_release_version,omitempty"`

	// Array of versions supported by this Looker instance
	// Read Only: true
	SupportedVersions []*APIVersionElement `json:"supported_versions"`
}

// Validate validates this Api version
func (m *APIVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIVersion) validateCurrentVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentVersion) { // not required
		return nil
	}

	if m.CurrentVersion != nil {
		if err := m.CurrentVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_version")
			}
			return err
		}
	}

	return nil
}

func (m *APIVersion) validateSupportedVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedVersions); i++ {
		if swag.IsZero(m.SupportedVersions[i]) { // not required
			continue
		}

		if m.SupportedVersions[i] != nil {
			if err := m.SupportedVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supported_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVersion) UnmarshalBinary(b []byte) error {
	var res APIVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}