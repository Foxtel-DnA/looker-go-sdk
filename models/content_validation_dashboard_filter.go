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

// ContentValidationDashboardFilter content validation dashboard filter
//
// swagger:model ContentValidationDashboardFilter
type ContentValidationDashboardFilter struct {

	// Id of Dashboard
	// Read Only: true
	DashboardID string `json:"dashboard_id,omitempty"`

	// Default value of filter
	DefaultValue string `json:"default_value,omitempty"`

	// Dimension of filter (required if type = field)
	Dimension string `json:"dimension,omitempty"`

	// Explore of filter (required if type = field)
	Explore string `json:"explore,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Model of filter (required if type = field)
	Model string `json:"model,omitempty"`

	// Name of filter
	Name string `json:"name,omitempty"`

	// Title of filter
	Title string `json:"title,omitempty"`

	// Type of filter: one of date, number, string, or field
	Type string `json:"type,omitempty"`
}

// Validate validates this content validation dashboard filter
func (m *ContentValidationDashboardFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this content validation dashboard filter based on the context it is used
func (m *ContentValidationDashboardFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboardID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidationDashboardFilter) contextValidateDashboardID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dashboard_id", "body", string(m.DashboardID)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidationDashboardFilter) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentValidationDashboardFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentValidationDashboardFilter) UnmarshalBinary(b []byte) error {
	var res ContentValidationDashboardFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}