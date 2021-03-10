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

// Session session
//
// swagger:model Session
type Session struct {

	// User's browser type
	// Read Only: true
	Browser string `json:"browser,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// City component of user location (derived from IP address)
	// Read Only: true
	City string `json:"city,omitempty"`

	// Country component of user location (derived from IP address)
	// Read Only: true
	Country string `json:"country,omitempty"`

	// Time when this session was initiated
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// Type of credentials used for logging in this session
	// Read Only: true
	CredentialsType string `json:"credentials_type,omitempty"`

	// Time when this session will expire
	// Read Only: true
	ExpiresAt string `json:"expires_at,omitempty"`

	// Time when this session was last extended by the user
	// Read Only: true
	ExtendedAt string `json:"extended_at,omitempty"`

	// Number of times this session was extended
	// Read Only: true
	ExtendedCount int64 `json:"extended_count,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// IP address of user when this session was initiated
	// Read Only: true
	IPAddress string `json:"ip_address,omitempty"`

	// User's Operating System
	// Read Only: true
	OperatingSystem string `json:"operating_system,omitempty"`

	// State component of user location (derived from IP address)
	// Read Only: true
	State string `json:"state,omitempty"`

	// Actual user in the case when this session represents one user sudo'ing as another
	// Read Only: true
	SudoUserID int64 `json:"sudo_user_id,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this session based on the context it is used
func (m *Session) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrowser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiresAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtendedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtendedCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingSystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSudoUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) contextValidateBrowser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "browser", "body", string(m.Browser)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Session) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "city", "body", string(m.City)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", string(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateCredentialsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "credentials_type", "body", string(m.CredentialsType)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateExpiresAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expires_at", "body", string(m.ExpiresAt)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateExtendedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "extended_at", "body", string(m.ExtendedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateExtendedCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "extended_count", "body", int64(m.ExtendedCount)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateIPAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ip_address", "body", string(m.IPAddress)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateOperatingSystem(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "operating_system", "body", string(m.OperatingSystem)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateSudoUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sudo_user_id", "body", int64(m.SudoUserID)); err != nil {
		return err
	}

	return nil
}

func (m *Session) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
