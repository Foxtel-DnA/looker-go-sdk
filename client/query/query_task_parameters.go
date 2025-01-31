// Code generated by go-swagger; DO NOT EDIT.

package query

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

// NewQueryTaskParams creates a new QueryTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryTaskParams() *QueryTaskParams {
	return &QueryTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryTaskParamsWithTimeout creates a new QueryTaskParams object
// with the ability to set a timeout on a request.
func NewQueryTaskParamsWithTimeout(timeout time.Duration) *QueryTaskParams {
	return &QueryTaskParams{
		timeout: timeout,
	}
}

// NewQueryTaskParamsWithContext creates a new QueryTaskParams object
// with the ability to set a context for a request.
func NewQueryTaskParamsWithContext(ctx context.Context) *QueryTaskParams {
	return &QueryTaskParams{
		Context: ctx,
	}
}

// NewQueryTaskParamsWithHTTPClient creates a new QueryTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryTaskParamsWithHTTPClient(client *http.Client) *QueryTaskParams {
	return &QueryTaskParams{
		HTTPClient: client,
	}
}

/* QueryTaskParams contains all the parameters to send to the API endpoint
   for the query task operation.

   Typically these are written to a http.Request.
*/
type QueryTaskParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* QueryTaskID.

	   ID of the Query Task
	*/
	QueryTaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTaskParams) WithDefaults() *QueryTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query task params
func (o *QueryTaskParams) WithTimeout(timeout time.Duration) *QueryTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query task params
func (o *QueryTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query task params
func (o *QueryTaskParams) WithContext(ctx context.Context) *QueryTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query task params
func (o *QueryTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query task params
func (o *QueryTaskParams) WithHTTPClient(client *http.Client) *QueryTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query task params
func (o *QueryTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the query task params
func (o *QueryTaskParams) WithFields(fields *string) *QueryTaskParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the query task params
func (o *QueryTaskParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithQueryTaskID adds the queryTaskID to the query task params
func (o *QueryTaskParams) WithQueryTaskID(queryTaskID string) *QueryTaskParams {
	o.SetQueryTaskID(queryTaskID)
	return o
}

// SetQueryTaskID adds the queryTaskId to the query task params
func (o *QueryTaskParams) SetQueryTaskID(queryTaskID string) {
	o.QueryTaskID = queryTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param query_task_id
	if err := r.SetPathParam("query_task_id", o.QueryTaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
