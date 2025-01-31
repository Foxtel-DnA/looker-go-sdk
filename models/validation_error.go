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
	"github.com/go-openapi/validate"
)

// ValidationError validation error
//
// swagger:model ValidationError
type ValidationError struct {

	// Documentation link
	// Required: true
	// Read Only: true
	// Format: uri
	DocumentationURL strfmt.URI `json:"documentation_url"`

	// Error detail array
	// Read Only: true
	Errors []*ValidationErrorDetail `json:"errors"`

	// Error details
	// Required: true
	// Read Only: true
	Message string `json:"message"`
}

// Validate validates this validation error
func (m *ValidationError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationError) validateDocumentationURL(formats strfmt.Registry) error {

	if err := validate.Required("documentation_url", "body", strfmt.URI(m.DocumentationURL)); err != nil {
		return err
	}

	if err := validate.FormatOf("documentation_url", "body", "uri", m.DocumentationURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ValidationError) validateErrors(formats strfmt.Registry) error {
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

func (m *ValidationError) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this validation error based on the context it is used
func (m *ValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentationURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationError) contextValidateDocumentationURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "documentation_url", "body", strfmt.URI(m.DocumentationURL)); err != nil {
		return err
	}

	return nil
}

func (m *ValidationError) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errors", "body", []*ValidationErrorDetail(m.Errors)); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ValidationError) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidationError) UnmarshalBinary(b []byte) error {
	var res ValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
