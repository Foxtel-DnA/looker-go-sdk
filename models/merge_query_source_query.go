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
)

// MergeQuerySourceQuery merge query source query
//
// swagger:model MergeQuerySourceQuery
type MergeQuerySourceQuery struct {

	// An array defining which fields of the source query are mapped onto fields of the merge query
	MergeFields []*MergeFields `json:"merge_fields"`

	// Display name
	Name string `json:"name,omitempty"`

	// Id of the query to merge
	QueryID int64 `json:"query_id,omitempty"`
}

// Validate validates this merge query source query
func (m *MergeQuerySourceQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMergeFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MergeQuerySourceQuery) validateMergeFields(formats strfmt.Registry) error {
	if swag.IsZero(m.MergeFields) { // not required
		return nil
	}

	for i := 0; i < len(m.MergeFields); i++ {
		if swag.IsZero(m.MergeFields[i]) { // not required
			continue
		}

		if m.MergeFields[i] != nil {
			if err := m.MergeFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("merge_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this merge query source query based on the context it is used
func (m *MergeQuerySourceQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMergeFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MergeQuerySourceQuery) contextValidateMergeFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MergeFields); i++ {

		if m.MergeFields[i] != nil {
			if err := m.MergeFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("merge_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MergeQuerySourceQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MergeQuerySourceQuery) UnmarshalBinary(b []byte) error {
	var res MergeQuerySourceQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
