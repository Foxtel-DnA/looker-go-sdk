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

// DashboardLayoutComponent dashboard layout component
//
// swagger:model DashboardLayoutComponent
type DashboardLayoutComponent struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Column
	Column int64 `json:"column,omitempty"`

	// Id Of Dashboard Element
	DashboardElementID string `json:"dashboard_element_id,omitempty"`

	// Id of Dashboard Layout
	DashboardLayoutID string `json:"dashboard_layout_id,omitempty"`

	// Whether or not the dashboard layout component is deleted
	// Read Only: true
	Deleted *bool `json:"deleted,omitempty"`

	// Dashboard element title, extracted from the Dashboard Element.
	// Read Only: true
	ElementTitle string `json:"element_title,omitempty"`

	// Whether or not the dashboard element title is displayed.
	// Read Only: true
	ElementTitleHidden *bool `json:"element_title_hidden,omitempty"`

	// Height
	Height int64 `json:"height,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Row
	Row int64 `json:"row,omitempty"`

	// Visualization type, extracted from a query's vis_config
	// Read Only: true
	VisType string `json:"vis_type,omitempty"`

	// Width
	Width int64 `json:"width,omitempty"`
}

// Validate validates this dashboard layout component
func (m *DashboardLayoutComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this dashboard layout component based on the context it is used
func (m *DashboardLayoutComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElementTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElementTitleHidden(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardLayoutComponent) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *DashboardLayoutComponent) contextValidateDeleted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *DashboardLayoutComponent) contextValidateElementTitle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "element_title", "body", string(m.ElementTitle)); err != nil {
		return err
	}

	return nil
}

func (m *DashboardLayoutComponent) contextValidateElementTitleHidden(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "element_title_hidden", "body", m.ElementTitleHidden); err != nil {
		return err
	}

	return nil
}

func (m *DashboardLayoutComponent) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DashboardLayoutComponent) contextValidateVisType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vis_type", "body", string(m.VisType)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardLayoutComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardLayoutComponent) UnmarshalBinary(b []byte) error {
	var res DashboardLayoutComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
