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

// NewDeleteUserLoginLockoutParams creates a new DeleteUserLoginLockoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserLoginLockoutParams() *DeleteUserLoginLockoutParams {
	return &DeleteUserLoginLockoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserLoginLockoutParamsWithTimeout creates a new DeleteUserLoginLockoutParams object
// with the ability to set a timeout on a request.
func NewDeleteUserLoginLockoutParamsWithTimeout(timeout time.Duration) *DeleteUserLoginLockoutParams {
	return &DeleteUserLoginLockoutParams{
		timeout: timeout,
	}
}

// NewDeleteUserLoginLockoutParamsWithContext creates a new DeleteUserLoginLockoutParams object
// with the ability to set a context for a request.
func NewDeleteUserLoginLockoutParamsWithContext(ctx context.Context) *DeleteUserLoginLockoutParams {
	return &DeleteUserLoginLockoutParams{
		Context: ctx,
	}
}

// NewDeleteUserLoginLockoutParamsWithHTTPClient creates a new DeleteUserLoginLockoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserLoginLockoutParamsWithHTTPClient(client *http.Client) *DeleteUserLoginLockoutParams {
	return &DeleteUserLoginLockoutParams{
		HTTPClient: client,
	}
}

/* DeleteUserLoginLockoutParams contains all the parameters to send to the API endpoint
   for the delete user login lockout operation.

   Typically these are written to a http.Request.
*/
type DeleteUserLoginLockoutParams struct {

	/* Key.

	   The key associated with the locked user
	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user login lockout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserLoginLockoutParams) WithDefaults() *DeleteUserLoginLockoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user login lockout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserLoginLockoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) WithTimeout(timeout time.Duration) *DeleteUserLoginLockoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) WithContext(ctx context.Context) *DeleteUserLoginLockoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) WithHTTPClient(client *http.Client) *DeleteUserLoginLockoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) WithKey(key string) *DeleteUserLoginLockoutParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete user login lockout params
func (o *DeleteUserLoginLockoutParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserLoginLockoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
