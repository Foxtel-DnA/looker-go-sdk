// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRoleGroupsParams creates a new RoleGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoleGroupsParams() *RoleGroupsParams {
	return &RoleGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoleGroupsParamsWithTimeout creates a new RoleGroupsParams object
// with the ability to set a timeout on a request.
func NewRoleGroupsParamsWithTimeout(timeout time.Duration) *RoleGroupsParams {
	return &RoleGroupsParams{
		timeout: timeout,
	}
}

// NewRoleGroupsParamsWithContext creates a new RoleGroupsParams object
// with the ability to set a context for a request.
func NewRoleGroupsParamsWithContext(ctx context.Context) *RoleGroupsParams {
	return &RoleGroupsParams{
		Context: ctx,
	}
}

// NewRoleGroupsParamsWithHTTPClient creates a new RoleGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoleGroupsParamsWithHTTPClient(client *http.Client) *RoleGroupsParams {
	return &RoleGroupsParams{
		HTTPClient: client,
	}
}

/* RoleGroupsParams contains all the parameters to send to the API endpoint
   for the role groups operation.

   Typically these are written to a http.Request.
*/
type RoleGroupsParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* RoleID.

	   id of role

	   Format: int64
	*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleGroupsParams) WithDefaults() *RoleGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the role groups params
func (o *RoleGroupsParams) WithTimeout(timeout time.Duration) *RoleGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role groups params
func (o *RoleGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role groups params
func (o *RoleGroupsParams) WithContext(ctx context.Context) *RoleGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role groups params
func (o *RoleGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role groups params
func (o *RoleGroupsParams) WithHTTPClient(client *http.Client) *RoleGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role groups params
func (o *RoleGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the role groups params
func (o *RoleGroupsParams) WithFields(fields *string) *RoleGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the role groups params
func (o *RoleGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithRoleID adds the roleID to the role groups params
func (o *RoleGroupsParams) WithRoleID(roleID int64) *RoleGroupsParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the role groups params
func (o *RoleGroupsParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RoleGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	// path param role_id
	if err := r.SetPathParam("role_id", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
