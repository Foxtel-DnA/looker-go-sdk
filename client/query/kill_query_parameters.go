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

// NewKillQueryParams creates a new KillQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKillQueryParams() *KillQueryParams {
	return &KillQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKillQueryParamsWithTimeout creates a new KillQueryParams object
// with the ability to set a timeout on a request.
func NewKillQueryParamsWithTimeout(timeout time.Duration) *KillQueryParams {
	return &KillQueryParams{
		timeout: timeout,
	}
}

// NewKillQueryParamsWithContext creates a new KillQueryParams object
// with the ability to set a context for a request.
func NewKillQueryParamsWithContext(ctx context.Context) *KillQueryParams {
	return &KillQueryParams{
		Context: ctx,
	}
}

// NewKillQueryParamsWithHTTPClient creates a new KillQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewKillQueryParamsWithHTTPClient(client *http.Client) *KillQueryParams {
	return &KillQueryParams{
		HTTPClient: client,
	}
}

/* KillQueryParams contains all the parameters to send to the API endpoint
   for the kill query operation.

   Typically these are written to a http.Request.
*/
type KillQueryParams struct {

	/* QueryTaskID.

	   Query task id.
	*/
	QueryTaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kill query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KillQueryParams) WithDefaults() *KillQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kill query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KillQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kill query params
func (o *KillQueryParams) WithTimeout(timeout time.Duration) *KillQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kill query params
func (o *KillQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kill query params
func (o *KillQueryParams) WithContext(ctx context.Context) *KillQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kill query params
func (o *KillQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kill query params
func (o *KillQueryParams) WithHTTPClient(client *http.Client) *KillQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kill query params
func (o *KillQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryTaskID adds the queryTaskID to the kill query params
func (o *KillQueryParams) WithQueryTaskID(queryTaskID string) *KillQueryParams {
	o.SetQueryTaskID(queryTaskID)
	return o
}

// SetQueryTaskID adds the queryTaskId to the kill query params
func (o *KillQueryParams) SetQueryTaskID(queryTaskID string) {
	o.QueryTaskID = queryTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *KillQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_task_id
	if err := r.SetPathParam("query_task_id", o.QueryTaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
