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

// Dialect dialect
//
// swagger:model Dialect
type Dialect struct {

	// Should SQL Runner snippets automatically be run
	// Read Only: true
	AutomaticallyRunSQLRunnerSnippets *bool `json:"automatically_run_sql_runner_snippets,omitempty"`

	// Array of names of the tests that can be run on a connection using this dialect
	// Read Only: true
	ConnectionTests []string `json:"connection_tests"`

	// Does the database have client SSL support settable through the JDBC string explicitly?
	// Read Only: true
	HasSslSupport *bool `json:"has_ssl_support,omitempty"`

	// The human-readable label of the connection
	// Read Only: true
	Label string `json:"label,omitempty"`

	// The name of the dialect
	// Read Only: true
	Name string `json:"name,omitempty"`

	// PDT distkey column
	// Read Only: true
	PersistentTableDistkey string `json:"persistent_table_distkey,omitempty"`

	// PDT index columns
	// Read Only: true
	PersistentTableIndexes string `json:"persistent_table_indexes,omitempty"`

	// PDT sortkey columns
	// Read Only: true
	PersistentTableSortkeys string `json:"persistent_table_sortkeys,omitempty"`

	// Whether the dialect supports query cost estimates
	// Read Only: true
	SupportsCostEstimate *bool `json:"supports_cost_estimate,omitempty"`

	// Is supported with the inducer (i.e. generate from sql)
	// Read Only: true
	SupportsInducer *bool `json:"supports_inducer,omitempty"`

	// Can multiple databases be accessed from a connection using this dialect
	// Read Only: true
	SupportsMultipleDatabases *bool `json:"supports_multiple_databases,omitempty"`

	// Whether the dialect supports allowing Looker to build persistent derived tables
	// Read Only: true
	SupportsPersistentDerivedTables *bool `json:"supports_persistent_derived_tables,omitempty"`

	// Suports streaming results
	// Read Only: true
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
}

// Validate validates this dialect
func (m *Dialect) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this dialect based on the context it is used
func (m *Dialect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutomaticallyRunSQLRunnerSnippets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionTests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHasSslSupport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersistentTableDistkey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersistentTableIndexes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersistentTableSortkeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsCostEstimate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsInducer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsMultipleDatabases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsPersistentDerivedTables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsStreaming(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dialect) contextValidateAutomaticallyRunSQLRunnerSnippets(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "automatically_run_sql_runner_snippets", "body", m.AutomaticallyRunSQLRunnerSnippets); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateConnectionTests(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connection_tests", "body", []string(m.ConnectionTests)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateHasSslSupport(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "has_ssl_support", "body", m.HasSslSupport); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "label", "body", string(m.Label)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidatePersistentTableDistkey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "persistent_table_distkey", "body", string(m.PersistentTableDistkey)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidatePersistentTableIndexes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "persistent_table_indexes", "body", string(m.PersistentTableIndexes)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidatePersistentTableSortkeys(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "persistent_table_sortkeys", "body", string(m.PersistentTableSortkeys)); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateSupportsCostEstimate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supports_cost_estimate", "body", m.SupportsCostEstimate); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateSupportsInducer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supports_inducer", "body", m.SupportsInducer); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateSupportsMultipleDatabases(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supports_multiple_databases", "body", m.SupportsMultipleDatabases); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateSupportsPersistentDerivedTables(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supports_persistent_derived_tables", "body", m.SupportsPersistentDerivedTables); err != nil {
		return err
	}

	return nil
}

func (m *Dialect) contextValidateSupportsStreaming(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supports_streaming", "body", m.SupportsStreaming); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dialect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dialect) UnmarshalBinary(b []byte) error {
	var res Dialect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
