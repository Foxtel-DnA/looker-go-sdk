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

// ResultMakerWithIDVisConfigAndDynamicFields result maker with Id vis config and dynamic fields
//
// swagger:model ResultMakerWithIdVisConfigAndDynamicFields
type ResultMakerWithIDVisConfigAndDynamicFields struct {

	// JSON string of dynamic field information.
	// Read Only: true
	DynamicFields string `json:"dynamic_fields,omitempty"`

	// array of items that can be filtered and information about them.
	// Read Only: true
	Filterables []*ResultMakerFilterables `json:"filterables"`

	// Unique Id.
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// ID of merge result if this is a merge_result.
	// Read Only: true
	MergeResultID string `json:"merge_result_id,omitempty"`

	// Query
	// Read Only: true
	Query *Query `json:"query,omitempty"`

	// ID of query if this is a query.
	// Read Only: true
	QueryID int64 `json:"query_id,omitempty"`

	// Sorts of the constituent Look, Query, or Merge Query
	// Read Only: true
	Sorts []string `json:"sorts"`

	// ID of SQL Query if this is a SQL Runner Query
	// Read Only: true
	SQLQueryID string `json:"sql_query_id,omitempty"`

	// Total of the constituent Look, Query, or Merge Query
	// Read Only: true
	Total *bool `json:"total,omitempty"`

	// Vis config of the constituent Query, or Merge Query.
	// Read Only: true
	VisConfig map[string]string `json:"vis_config,omitempty"`
}

// Validate validates this result maker with Id vis config and dynamic fields
func (m *ResultMakerWithIDVisConfigAndDynamicFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterables(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) validateFilterables(formats strfmt.Registry) error {
	if swag.IsZero(m.Filterables) { // not required
		return nil
	}

	for i := 0; i < len(m.Filterables); i++ {
		if swag.IsZero(m.Filterables[i]) { // not required
			continue
		}

		if m.Filterables[i] != nil {
			if err := m.Filterables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filterables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this result maker with Id vis config and dynamic fields based on the context it is used
func (m *ResultMakerWithIDVisConfigAndDynamicFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDynamicFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMergeResultID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLQueryID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateDynamicFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dynamic_fields", "body", string(m.DynamicFields)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateFilterables(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "filterables", "body", []*ResultMakerFilterables(m.Filterables)); err != nil {
		return err
	}

	for i := 0; i < len(m.Filterables); i++ {

		if m.Filterables[i] != nil {
			if err := m.Filterables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filterables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateMergeResultID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "merge_result_id", "body", string(m.MergeResultID)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {
		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateQueryID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "query_id", "body", int64(m.QueryID)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateSorts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sorts", "body", []string(m.Sorts)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateSQLQueryID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sql_query_id", "body", string(m.SQLQueryID)); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *ResultMakerWithIDVisConfigAndDynamicFields) contextValidateVisConfig(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ResultMakerWithIDVisConfigAndDynamicFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultMakerWithIDVisConfigAndDynamicFields) UnmarshalBinary(b []byte) error {
	var res ResultMakerWithIDVisConfigAndDynamicFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
