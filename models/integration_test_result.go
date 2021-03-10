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

// IntegrationTestResult integration test result
//
// swagger:model IntegrationTestResult
type IntegrationTestResult struct {

	// An array of connection test result for delegate oauth actions.
	// Read Only: true
	DelegateOauthResult []*DelegateOauthTest `json:"delegate_oauth_result"`

	// A message representing the results of the test.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Whether or not the test was successful
	// Read Only: true
	Success *bool `json:"success,omitempty"`
}

// Validate validates this integration test result
func (m *IntegrationTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelegateOauthResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationTestResult) validateDelegateOauthResult(formats strfmt.Registry) error {
	if swag.IsZero(m.DelegateOauthResult) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegateOauthResult); i++ {
		if swag.IsZero(m.DelegateOauthResult[i]) { // not required
			continue
		}

		if m.DelegateOauthResult[i] != nil {
			if err := m.DelegateOauthResult[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegate_oauth_result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this integration test result based on the context it is used
func (m *IntegrationTestResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDelegateOauthResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationTestResult) contextValidateDelegateOauthResult(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "delegate_oauth_result", "body", []*DelegateOauthTest(m.DelegateOauthResult)); err != nil {
		return err
	}

	for i := 0; i < len(m.DelegateOauthResult); i++ {

		if m.DelegateOauthResult[i] != nil {
			if err := m.DelegateOauthResult[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegate_oauth_result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IntegrationTestResult) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationTestResult) contextValidateSuccess(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationTestResult) UnmarshalBinary(b []byte) error {
	var res IntegrationTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
