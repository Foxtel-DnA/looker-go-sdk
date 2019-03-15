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

// LookmlModel lookml model
// swagger:model LookmlModel
type LookmlModel struct {

	// Array of names of connections this model is allowed to use
	AllowedDbConnectionNames []string `json:"allowed_db_connection_names"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Array of explores (if has_content)
	// Read Only: true
	Explores []*LookmlModelNavExplore `json:"explores"`

	// Does this model declaration have have lookml content?
	// Read Only: true
	HasContent *bool `json:"has_content,omitempty"`

	// UI-friendly name for this model
	// Read Only: true
	Label string `json:"label,omitempty"`

	// Name of the model. Also used as the unique identifier
	Name string `json:"name,omitempty"`

	// Name of project containing the model
	ProjectName string `json:"project_name,omitempty"`

	// Is this model allowed to use all current and future connections
	UnlimitedDbConnections bool `json:"unlimited_db_connections,omitempty"`
}

// Validate validates this lookml model
func (m *LookmlModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExplores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModel) validateExplores(formats strfmt.Registry) error {

	if swag.IsZero(m.Explores) { // not required
		return nil
	}

	for i := 0; i < len(m.Explores); i++ {
		if swag.IsZero(m.Explores[i]) { // not required
			continue
		}

		if m.Explores[i] != nil {
			if err := m.Explores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("explores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModel) UnmarshalBinary(b []byte) error {
	var res LookmlModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}