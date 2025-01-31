// Code generated by go-swagger; DO NOT EDIT.

package project

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
)

// NewValidateProjectParams creates a new ValidateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateProjectParams() *ValidateProjectParams {
	return &ValidateProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateProjectParamsWithTimeout creates a new ValidateProjectParams object
// with the ability to set a timeout on a request.
func NewValidateProjectParamsWithTimeout(timeout time.Duration) *ValidateProjectParams {
	return &ValidateProjectParams{
		timeout: timeout,
	}
}

// NewValidateProjectParamsWithContext creates a new ValidateProjectParams object
// with the ability to set a context for a request.
func NewValidateProjectParamsWithContext(ctx context.Context) *ValidateProjectParams {
	return &ValidateProjectParams{
		Context: ctx,
	}
}

// NewValidateProjectParamsWithHTTPClient creates a new ValidateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateProjectParamsWithHTTPClient(client *http.Client) *ValidateProjectParams {
	return &ValidateProjectParams{
		HTTPClient: client,
	}
}

/* ValidateProjectParams contains all the parameters to send to the API endpoint
   for the validate project operation.

   Typically these are written to a http.Request.
*/
type ValidateProjectParams struct {

	/* Fields.

	   Requested fields
	*/
	Fields *string

	/* ProjectID.

	   Project Id
	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateProjectParams) WithDefaults() *ValidateProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate project params
func (o *ValidateProjectParams) WithTimeout(timeout time.Duration) *ValidateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate project params
func (o *ValidateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate project params
func (o *ValidateProjectParams) WithContext(ctx context.Context) *ValidateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate project params
func (o *ValidateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate project params
func (o *ValidateProjectParams) WithHTTPClient(client *http.Client) *ValidateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate project params
func (o *ValidateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the validate project params
func (o *ValidateProjectParams) WithFields(fields *string) *ValidateProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the validate project params
func (o *ValidateProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectID adds the projectID to the validate project params
func (o *ValidateProjectParams) WithProjectID(projectID string) *ValidateProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the validate project params
func (o *ValidateProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
