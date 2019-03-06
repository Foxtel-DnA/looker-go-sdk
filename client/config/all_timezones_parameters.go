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

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllTimezonesParams creates a new AllTimezonesParams object
// with the default values initialized.
func NewAllTimezonesParams() *AllTimezonesParams {

	return &AllTimezonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllTimezonesParamsWithTimeout creates a new AllTimezonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllTimezonesParamsWithTimeout(timeout time.Duration) *AllTimezonesParams {

	return &AllTimezonesParams{

		timeout: timeout,
	}
}

// NewAllTimezonesParamsWithContext creates a new AllTimezonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllTimezonesParamsWithContext(ctx context.Context) *AllTimezonesParams {

	return &AllTimezonesParams{

		Context: ctx,
	}
}

// NewAllTimezonesParamsWithHTTPClient creates a new AllTimezonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllTimezonesParamsWithHTTPClient(client *http.Client) *AllTimezonesParams {

	return &AllTimezonesParams{
		HTTPClient: client,
	}
}

/*AllTimezonesParams contains all the parameters to send to the API endpoint
for the all timezones operation typically these are written to a http.Request
*/
type AllTimezonesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all timezones params
func (o *AllTimezonesParams) WithTimeout(timeout time.Duration) *AllTimezonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all timezones params
func (o *AllTimezonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all timezones params
func (o *AllTimezonesParams) WithContext(ctx context.Context) *AllTimezonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all timezones params
func (o *AllTimezonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all timezones params
func (o *AllTimezonesParams) WithHTTPClient(client *http.Client) *AllTimezonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all timezones params
func (o *AllTimezonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AllTimezonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
