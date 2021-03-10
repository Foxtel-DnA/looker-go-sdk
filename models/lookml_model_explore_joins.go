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

// LookmlModelExploreJoins lookml model explore joins
//
// swagger:model LookmlModelExploreJoins
type LookmlModelExploreJoins struct {

	// Fields referenced by the join
	// Read Only: true
	DependentFields []string `json:"dependent_fields"`

	// Fields of the joined view to pull into this explore
	// Read Only: true
	Fields []string `json:"fields"`

	// Name of the dimension in this explore whose value is in the primary key of the joined view
	// Read Only: true
	ForeignKey string `json:"foreign_key,omitempty"`

	// Name of view to join
	// Read Only: true
	From string `json:"from,omitempty"`

	// Name of this join (and name of the view to join)
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Specifies whether all queries must use an outer join
	// Read Only: true
	OuterOnly *bool `json:"outer_only,omitempty"`

	// many_to_one, one_to_one, one_to_many, many_to_many
	// Read Only: true
	Relationship string `json:"relationship,omitempty"`

	// Names of joins that must always be included in SQL queries
	// Read Only: true
	RequiredJoins []string `json:"required_joins"`

	// SQL expression that produces a foreign key
	// Read Only: true
	SQLForeignKey string `json:"sql_foreign_key,omitempty"`

	// SQL ON expression describing the join condition
	// Read Only: true
	SQLOn string `json:"sql_on,omitempty"`

	// SQL table name to join
	// Read Only: true
	SQLTableName string `json:"sql_table_name,omitempty"`

	// The join type: left_outer, full_outer, inner, or cross
	// Read Only: true
	Type string `json:"type,omitempty"`

	// Label to display in UI selectors
	// Read Only: true
	ViewLabel string `json:"view_label,omitempty"`
}

// Validate validates this lookml model explore joins
func (m *LookmlModelExploreJoins) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this lookml model explore joins based on the context it is used
func (m *LookmlModelExploreJoins) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependentFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForeignKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOuterOnly(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredJoins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLForeignKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLTableName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModelExploreJoins) contextValidateDependentFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dependent_fields", "body", []string(m.DependentFields)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "fields", "body", []string(m.Fields)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateForeignKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "foreign_key", "body", string(m.ForeignKey)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "from", "body", string(m.From)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateOuterOnly(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "outer_only", "body", m.OuterOnly); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateRelationship(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "relationship", "body", string(m.Relationship)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateRequiredJoins(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "required_joins", "body", []string(m.RequiredJoins)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateSQLForeignKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sql_foreign_key", "body", string(m.SQLForeignKey)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateSQLOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sql_on", "body", string(m.SQLOn)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateSQLTableName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sql_table_name", "body", string(m.SQLTableName)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreJoins) contextValidateViewLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "view_label", "body", string(m.ViewLabel)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreJoins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreJoins) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreJoins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
