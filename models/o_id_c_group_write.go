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

// OIDCGroupWrite o ID c group write
//
// swagger:model OIDCGroupWrite
type OIDCGroupWrite struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Unique Id
	ID int64 `json:"id,omitempty"`

	// Unique Id of group in Looker
	// Read Only: true
	LookerGroupID int64 `json:"looker_group_id,omitempty"`

	// Name of group in Looker
	LookerGroupName string `json:"looker_group_name,omitempty"`

	// Name of group in OIDC
	Name string `json:"name,omitempty"`

	// Looker Role Ids
	RoleIds []int64 `json:"role_ids"`
}

// Validate validates this o ID c group write
func (m *OIDCGroupWrite) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this o ID c group write based on the context it is used
func (m *OIDCGroupWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLookerGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OIDCGroupWrite) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *OIDCGroupWrite) contextValidateLookerGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "looker_group_id", "body", int64(m.LookerGroupID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OIDCGroupWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCGroupWrite) UnmarshalBinary(b []byte) error {
	var res OIDCGroupWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
