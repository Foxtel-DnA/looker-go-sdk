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

// APISession Api session
//
// swagger:model ApiSession
type APISession struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// The id of the actual user in the case when this session represents one user sudo'ing as another
	// Read Only: true
	SudoUserID int64 `json:"sudo_user_id,omitempty"`

	// The id of active workspace for this session
	WorkspaceID string `json:"workspace_id,omitempty"`
}

// Validate validates this Api session
func (m *APISession) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this Api session based on the context it is used
func (m *APISession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSudoUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISession) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *APISession) contextValidateSudoUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sudo_user_id", "body", int64(m.SudoUserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APISession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISession) UnmarshalBinary(b []byte) error {
	var res APISession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
