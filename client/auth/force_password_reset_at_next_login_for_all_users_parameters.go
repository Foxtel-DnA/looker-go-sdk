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

// NewForcePasswordResetAtNextLoginForAllUsersParams creates a new ForcePasswordResetAtNextLoginForAllUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForcePasswordResetAtNextLoginForAllUsersParams() *ForcePasswordResetAtNextLoginForAllUsersParams {
	return &ForcePasswordResetAtNextLoginForAllUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForcePasswordResetAtNextLoginForAllUsersParamsWithTimeout creates a new ForcePasswordResetAtNextLoginForAllUsersParams object
// with the ability to set a timeout on a request.
func NewForcePasswordResetAtNextLoginForAllUsersParamsWithTimeout(timeout time.Duration) *ForcePasswordResetAtNextLoginForAllUsersParams {
	return &ForcePasswordResetAtNextLoginForAllUsersParams{
		timeout: timeout,
	}
}

// NewForcePasswordResetAtNextLoginForAllUsersParamsWithContext creates a new ForcePasswordResetAtNextLoginForAllUsersParams object
// with the ability to set a context for a request.
func NewForcePasswordResetAtNextLoginForAllUsersParamsWithContext(ctx context.Context) *ForcePasswordResetAtNextLoginForAllUsersParams {
	return &ForcePasswordResetAtNextLoginForAllUsersParams{
		Context: ctx,
	}
}

// NewForcePasswordResetAtNextLoginForAllUsersParamsWithHTTPClient creates a new ForcePasswordResetAtNextLoginForAllUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewForcePasswordResetAtNextLoginForAllUsersParamsWithHTTPClient(client *http.Client) *ForcePasswordResetAtNextLoginForAllUsersParams {
	return &ForcePasswordResetAtNextLoginForAllUsersParams{
		HTTPClient: client,
	}
}

/* ForcePasswordResetAtNextLoginForAllUsersParams contains all the parameters to send to the API endpoint
   for the force password reset at next login for all users operation.

   Typically these are written to a http.Request.
*/
type ForcePasswordResetAtNextLoginForAllUsersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the force password reset at next login for all users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) WithDefaults() *ForcePasswordResetAtNextLoginForAllUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the force password reset at next login for all users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) WithTimeout(timeout time.Duration) *ForcePasswordResetAtNextLoginForAllUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) WithContext(ctx context.Context) *ForcePasswordResetAtNextLoginForAllUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) WithHTTPClient(client *http.Client) *ForcePasswordResetAtNextLoginForAllUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the force password reset at next login for all users params
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ForcePasswordResetAtNextLoginForAllUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
