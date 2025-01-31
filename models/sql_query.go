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

// SQLQuery Sql query
//
// swagger:model SqlQuery
type SQLQuery struct {

	// Maximum number of rows this query will display on the SQL Runner page
	// Read Only: true
	BrowserLimit int64 `json:"browser_limit,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Connection this query uses
	// Read Only: true
	Connection *DBConnectionBase `json:"connection,omitempty"`

	// User who created this SQL query
	// Read Only: true
	Creator *UserPublic `json:"creator,omitempty"`

	// Explore page URL for this SQL query
	// Read Only: true
	ExploreURL string `json:"explore_url,omitempty"`

	// The most recent time this query was run
	// Read Only: true
	LastRunAt string `json:"last_run_at,omitempty"`

	// Number of seconds this query took to run the most recent time it was run
	// Read Only: true
	LastRuntime float32 `json:"last_runtime,omitempty"`

	// Model name this query uses
	// Read Only: true
	ModelName string `json:"model_name,omitempty"`

	// Should this query be rendered as plain text
	// Read Only: true
	Plaintext *bool `json:"plaintext,omitempty"`

	// ID of the ResultMakerLookup entry.
	ResultMakerID int64 `json:"result_maker_id,omitempty"`

	// Number of times this query has been run
	// Read Only: true
	RunCount int64 `json:"run_count,omitempty"`

	// The identifier of the SQL query
	// Read Only: true
	Slug string `json:"slug,omitempty"`

	// SQL query text
	// Read Only: true
	SQL string `json:"sql,omitempty"`

	// Visualization configuration properties. These properties are typically opaque and differ based on the type of visualization used. There is no specified set of allowed keys. The values can be any type supported by JSON. A "type" key with a string value is often present, and is used by Looker to determine which visualization to present. Visualizations ignore unknown vis_config properties.
	VisConfig map[string]string `json:"vis_config,omitempty"`
}

// Validate validates this Sql query
func (m *SQLQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLQuery) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *SQLQuery) validateCreator(formats strfmt.Registry) error {
	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {
		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Sql query based on the context it is used
func (m *SQLQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrowserLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExploreURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastRunAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastRuntime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModelName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlaintext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlug(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLQuery) contextValidateBrowserLimit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "browser_limit", "body", int64(m.BrowserLimit)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *SQLQuery) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *SQLQuery) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if m.Creator != nil {
		if err := m.Creator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

func (m *SQLQuery) contextValidateExploreURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "explore_url", "body", string(m.ExploreURL)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateLastRunAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_run_at", "body", string(m.LastRunAt)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateLastRuntime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_runtime", "body", float32(m.LastRuntime)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateModelName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "model_name", "body", string(m.ModelName)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidatePlaintext(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "plaintext", "body", m.Plaintext); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateRunCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "run_count", "body", int64(m.RunCount)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateSlug(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "slug", "body", string(m.Slug)); err != nil {
		return err
	}

	return nil
}

func (m *SQLQuery) contextValidateSQL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sql", "body", string(m.SQL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLQuery) UnmarshalBinary(b []byte) error {
	var res SQLQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
