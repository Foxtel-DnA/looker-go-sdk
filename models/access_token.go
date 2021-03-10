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

// AccessToken access token
//
// swagger:model AccessToken
type AccessToken struct {

	// Access Token used for API calls
	// Read Only: true
	AccessToken string `json:"access_token,omitempty"`

	// Number of seconds before the token expires
	// Read Only: true
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// Type of Token
	// Read Only: true
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this access token
func (m *AccessToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this access token based on the context it is used
func (m *AccessToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiresIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTokenType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessToken) contextValidateAccessToken(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "access_token", "body", string(m.AccessToken)); err != nil {
		return err
	}

	return nil
}

func (m *AccessToken) contextValidateExpiresIn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expires_in", "body", int64(m.ExpiresIn)); err != nil {
		return err
	}

	return nil
}

func (m *AccessToken) contextValidateTokenType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "token_type", "body", string(m.TokenType)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessToken) UnmarshalBinary(b []byte) error {
	var res AccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
