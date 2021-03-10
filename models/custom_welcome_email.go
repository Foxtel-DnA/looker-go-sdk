// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomWelcomeEmail custom welcome email
//
// swagger:model CustomWelcomeEmail
type CustomWelcomeEmail struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// The HTML to use as custom content for welcome emails. Script elements and other potentially dangerous markup will be removed.
	Content string `json:"content,omitempty"`

	// If true, custom email content will replace the default body of welcome emails
	Enabled bool `json:"enabled,omitempty"`

	// The text to appear in the header line of the email body.
	Header string `json:"header,omitempty"`

	// The text to appear in the email subject line.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this custom welcome email
func (m *CustomWelcomeEmail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this custom welcome email based on the context it is used
func (m *CustomWelcomeEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomWelcomeEmail) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CustomWelcomeEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomWelcomeEmail) UnmarshalBinary(b []byte) error {
	var res CustomWelcomeEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
