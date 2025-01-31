// Code generated by go-swagger; DO NOT EDIT.

package folder

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

// NewAllFoldersParams creates a new AllFoldersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllFoldersParams() *AllFoldersParams {
	return &AllFoldersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllFoldersParamsWithTimeout creates a new AllFoldersParams object
// with the ability to set a timeout on a request.
func NewAllFoldersParamsWithTimeout(timeout time.Duration) *AllFoldersParams {
	return &AllFoldersParams{
		timeout: timeout,
	}
}

// NewAllFoldersParamsWithContext creates a new AllFoldersParams object
// with the ability to set a context for a request.
func NewAllFoldersParamsWithContext(ctx context.Context) *AllFoldersParams {
	return &AllFoldersParams{
		Context: ctx,
	}
}

// NewAllFoldersParamsWithHTTPClient creates a new AllFoldersParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllFoldersParamsWithHTTPClient(client *http.Client) *AllFoldersParams {
	return &AllFoldersParams{
		HTTPClient: client,
	}
}

/* AllFoldersParams contains all the parameters to send to the API endpoint
   for the all folders operation.

   Typically these are written to a http.Request.
*/
type AllFoldersParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllFoldersParams) WithDefaults() *AllFoldersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllFoldersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all folders params
func (o *AllFoldersParams) WithTimeout(timeout time.Duration) *AllFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all folders params
func (o *AllFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all folders params
func (o *AllFoldersParams) WithContext(ctx context.Context) *AllFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all folders params
func (o *AllFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all folders params
func (o *AllFoldersParams) WithHTTPClient(client *http.Client) *AllFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all folders params
func (o *AllFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all folders params
func (o *AllFoldersParams) WithFields(fields *string) *AllFoldersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all folders params
func (o *AllFoldersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
