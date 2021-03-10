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

// LookmlModelExploreFieldTimeInterval lookml model explore field time interval
//
// swagger:model LookmlModelExploreFieldTimeInterval
type LookmlModelExploreFieldTimeInterval struct {

	// The number of intervals this field represents a grouping of.
	// Read Only: true
	Count int64 `json:"count,omitempty"`

	// The type of time interval this field represents a grouping of. Valid values are: "day", "hour", "minute", "second", "millisecond", "microsecond", "week", "month", "quarter", "year".
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this lookml model explore field time interval
func (m *LookmlModelExploreFieldTimeInterval) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this lookml model explore field time interval based on the context it is used
func (m *LookmlModelExploreFieldTimeInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCount(ctx, formats); err != nil {
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

func (m *LookmlModelExploreFieldTimeInterval) contextValidateCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count", "body", int64(m.Count)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreFieldTimeInterval) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreFieldTimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreFieldTimeInterval) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreFieldTimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
