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

// ContentValidation content validation
//
// swagger:model ContentValidation
type ContentValidation struct {

	// Duration of content validation in seconds
	// Read Only: true
	ComputationTime float32 `json:"computation_time,omitempty"`

	// A list of content errors
	// Read Only: true
	ContentWithErrors []*ContentValidatorError `json:"content_with_errors"`

	// The number of alerts validated
	// Read Only: true
	TotalAlertsValidated int64 `json:"total_alerts_validated,omitempty"`

	// The number of dashboard elements validated
	// Read Only: true
	TotalDashboardElementsValidated int64 `json:"total_dashboard_elements_validated,omitempty"`

	// The number of dashboard filters validated
	// Read Only: true
	TotalDashboardFiltersValidated int64 `json:"total_dashboard_filters_validated,omitempty"`

	// The number of explores used across all content validated
	// Read Only: true
	TotalExploresValidated int64 `json:"total_explores_validated,omitempty"`

	// The number of looks validated
	// Read Only: true
	TotalLooksValidated int64 `json:"total_looks_validated,omitempty"`

	// The number of scheduled plans validated
	// Read Only: true
	TotalScheduledPlansValidated int64 `json:"total_scheduled_plans_validated,omitempty"`
}

// Validate validates this content validation
func (m *ContentValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentWithErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidation) validateContentWithErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentWithErrors) { // not required
		return nil
	}

	for i := 0; i < len(m.ContentWithErrors); i++ {
		if swag.IsZero(m.ContentWithErrors[i]) { // not required
			continue
		}

		if m.ContentWithErrors[i] != nil {
			if err := m.ContentWithErrors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content_with_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this content validation based on the context it is used
func (m *ContentValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputationTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentWithErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalAlertsValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalDashboardElementsValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalDashboardFiltersValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalExploresValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalLooksValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalScheduledPlansValidated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidation) contextValidateComputationTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computation_time", "body", float32(m.ComputationTime)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateContentWithErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content_with_errors", "body", []*ContentValidatorError(m.ContentWithErrors)); err != nil {
		return err
	}

	for i := 0; i < len(m.ContentWithErrors); i++ {

		if m.ContentWithErrors[i] != nil {
			if err := m.ContentWithErrors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content_with_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentValidation) contextValidateTotalAlertsValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_alerts_validated", "body", int64(m.TotalAlertsValidated)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateTotalDashboardElementsValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_dashboard_elements_validated", "body", int64(m.TotalDashboardElementsValidated)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateTotalDashboardFiltersValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_dashboard_filters_validated", "body", int64(m.TotalDashboardFiltersValidated)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateTotalExploresValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_explores_validated", "body", int64(m.TotalExploresValidated)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateTotalLooksValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_looks_validated", "body", int64(m.TotalLooksValidated)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidation) contextValidateTotalScheduledPlansValidated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_scheduled_plans_validated", "body", int64(m.TotalScheduledPlansValidated)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentValidation) UnmarshalBinary(b []byte) error {
	var res ContentValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
