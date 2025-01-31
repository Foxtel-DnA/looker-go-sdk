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

// InternalHelpResourcesContent internal help resources content
//
// swagger:model InternalHelpResourcesContent
type InternalHelpResourcesContent struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Content to be displayed in the internal help resources page/modal
	MarkdownContent string `json:"markdown_content,omitempty"`

	// Text to display in the help menu item which will display the internal help resources
	OrganizationName string `json:"organization_name,omitempty"`
}

// Validate validates this internal help resources content
func (m *InternalHelpResourcesContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this internal help resources content based on the context it is used
func (m *InternalHelpResourcesContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternalHelpResourcesContent) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *InternalHelpResourcesContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalHelpResourcesContent) UnmarshalBinary(b []byte) error {
	var res InternalHelpResourcesContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
