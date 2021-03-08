// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewPasswordConfigParams creates a new PasswordConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPasswordConfigParams() *PasswordConfigParams {
	return &PasswordConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPasswordConfigParamsWithTimeout creates a new PasswordConfigParams object
// with the ability to set a timeout on a request.
func NewPasswordConfigParamsWithTimeout(timeout time.Duration) *PasswordConfigParams {
	return &PasswordConfigParams{
		timeout: timeout,
	}
}

// NewPasswordConfigParamsWithContext creates a new PasswordConfigParams object
// with the ability to set a context for a request.
func NewPasswordConfigParamsWithContext(ctx context.Context) *PasswordConfigParams {
	return &PasswordConfigParams{
		Context: ctx,
	}
}

// NewPasswordConfigParamsWithHTTPClient creates a new PasswordConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPasswordConfigParamsWithHTTPClient(client *http.Client) *PasswordConfigParams {
	return &PasswordConfigParams{
		HTTPClient: client,
	}
}

/* PasswordConfigParams contains all the parameters to send to the API endpoint
   for the password config operation.

   Typically these are written to a http.Request.
*/
type PasswordConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the password config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PasswordConfigParams) WithDefaults() *PasswordConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the password config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PasswordConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the password config params
func (o *PasswordConfigParams) WithTimeout(timeout time.Duration) *PasswordConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the password config params
func (o *PasswordConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the password config params
func (o *PasswordConfigParams) WithContext(ctx context.Context) *PasswordConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the password config params
func (o *PasswordConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the password config params
func (o *PasswordConfigParams) WithHTTPClient(client *http.Client) *PasswordConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the password config params
func (o *PasswordConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PasswordConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}