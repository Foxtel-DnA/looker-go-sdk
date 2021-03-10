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

// NewAllLegacyFeaturesParams creates a new AllLegacyFeaturesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllLegacyFeaturesParams() *AllLegacyFeaturesParams {
	return &AllLegacyFeaturesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllLegacyFeaturesParamsWithTimeout creates a new AllLegacyFeaturesParams object
// with the ability to set a timeout on a request.
func NewAllLegacyFeaturesParamsWithTimeout(timeout time.Duration) *AllLegacyFeaturesParams {
	return &AllLegacyFeaturesParams{
		timeout: timeout,
	}
}

// NewAllLegacyFeaturesParamsWithContext creates a new AllLegacyFeaturesParams object
// with the ability to set a context for a request.
func NewAllLegacyFeaturesParamsWithContext(ctx context.Context) *AllLegacyFeaturesParams {
	return &AllLegacyFeaturesParams{
		Context: ctx,
	}
}

// NewAllLegacyFeaturesParamsWithHTTPClient creates a new AllLegacyFeaturesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllLegacyFeaturesParamsWithHTTPClient(client *http.Client) *AllLegacyFeaturesParams {
	return &AllLegacyFeaturesParams{
		HTTPClient: client,
	}
}

/* AllLegacyFeaturesParams contains all the parameters to send to the API endpoint
   for the all legacy features operation.

   Typically these are written to a http.Request.
*/
type AllLegacyFeaturesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all legacy features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllLegacyFeaturesParams) WithDefaults() *AllLegacyFeaturesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all legacy features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllLegacyFeaturesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all legacy features params
func (o *AllLegacyFeaturesParams) WithTimeout(timeout time.Duration) *AllLegacyFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all legacy features params
func (o *AllLegacyFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all legacy features params
func (o *AllLegacyFeaturesParams) WithContext(ctx context.Context) *AllLegacyFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all legacy features params
func (o *AllLegacyFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all legacy features params
func (o *AllLegacyFeaturesParams) WithHTTPClient(client *http.Client) *AllLegacyFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all legacy features params
func (o *AllLegacyFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AllLegacyFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
