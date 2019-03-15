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

// ProjectValidationCache project validation cache
// swagger:model ProjectValidationCache
type ProjectValidationCache struct {

	// Duration of project validation in seconds
	// Read Only: true
	ComputationTime float32 `json:"computation_time,omitempty"`

	// A list of project errors
	// Read Only: true
	Errors []*ProjectError `json:"errors"`

	// A list of models which were not fully validated
	// Read Only: true
	ModelsNotValidated []*ModelsNotValidated `json:"models_not_validated"`

	// A hash value computed from the project's current state
	// Read Only: true
	ProjectDigest string `json:"project_digest,omitempty"`

	// If true, the cached project validation results are no longer accurate because the project has changed since the cached results were calculated
	// Read Only: true
	Stale *bool `json:"stale,omitempty"`
}

// Validate validates this project validation cache
func (m *ProjectValidationCache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelsNotValidated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectValidationCache) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectValidationCache) validateModelsNotValidated(formats strfmt.Registry) error {

	if swag.IsZero(m.ModelsNotValidated) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelsNotValidated); i++ {
		if swag.IsZero(m.ModelsNotValidated[i]) { // not required
			continue
		}

		if m.ModelsNotValidated[i] != nil {
			if err := m.ModelsNotValidated[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("models_not_validated" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectValidationCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectValidationCache) UnmarshalBinary(b []byte) error {
	var res ProjectValidationCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}