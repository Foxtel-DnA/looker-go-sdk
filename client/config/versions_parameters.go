// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewVersionsParams creates a new VersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionsParams() *VersionsParams {
	return &VersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionsParamsWithTimeout creates a new VersionsParams object
// with the ability to set a timeout on a request.
func NewVersionsParamsWithTimeout(timeout time.Duration) *VersionsParams {
	return &VersionsParams{
		timeout: timeout,
	}
}

// NewVersionsParamsWithContext creates a new VersionsParams object
// with the ability to set a context for a request.
func NewVersionsParamsWithContext(ctx context.Context) *VersionsParams {
	return &VersionsParams{
		Context: ctx,
	}
}

// NewVersionsParamsWithHTTPClient creates a new VersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionsParamsWithHTTPClient(client *http.Client) *VersionsParams {
	return &VersionsParams{
		HTTPClient: client,
	}
}

/* VersionsParams contains all the parameters to send to the API endpoint
   for the versions operation.

   Typically these are written to a http.Request.
*/
type VersionsParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionsParams) WithDefaults() *VersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the versions params
func (o *VersionsParams) WithTimeout(timeout time.Duration) *VersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the versions params
func (o *VersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the versions params
func (o *VersionsParams) WithContext(ctx context.Context) *VersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the versions params
func (o *VersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the versions params
func (o *VersionsParams) WithHTTPClient(client *http.Client) *VersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the versions params
func (o *VersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the versions params
func (o *VersionsParams) WithFields(fields *string) *VersionsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the versions params
func (o *VersionsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *VersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
