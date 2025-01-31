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

// BackupConfiguration backup configuration
//
// swagger:model BackupConfiguration
type BackupConfiguration struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Name of bucket for custom-s3 backups
	CustomS3Bucket string `json:"custom_s3_bucket,omitempty"`

	// Name of region where the bucket is located
	CustomS3BucketRegion string `json:"custom_s3_bucket_region,omitempty"`

	// (Write-Only) AWS S3 key used for custom-s3 backups
	CustomS3Key string `json:"custom_s3_key,omitempty"`

	// (Write-Only) AWS S3 secret used for custom-s3 backups
	CustomS3Secret string `json:"custom_s3_secret,omitempty"`

	// Type of backup: looker-s3 or custom-s3
	Type string `json:"type,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this backup configuration
func (m *BackupConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupConfiguration) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup configuration based on the context it is used
func (m *BackupConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
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

func (m *BackupConfiguration) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupConfiguration) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupConfiguration) UnmarshalBinary(b []byte) error {
	var res BackupConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
