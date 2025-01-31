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

// NewLdapConfigParams creates a new LdapConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapConfigParams() *LdapConfigParams {
	return &LdapConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapConfigParamsWithTimeout creates a new LdapConfigParams object
// with the ability to set a timeout on a request.
func NewLdapConfigParamsWithTimeout(timeout time.Duration) *LdapConfigParams {
	return &LdapConfigParams{
		timeout: timeout,
	}
}

// NewLdapConfigParamsWithContext creates a new LdapConfigParams object
// with the ability to set a context for a request.
func NewLdapConfigParamsWithContext(ctx context.Context) *LdapConfigParams {
	return &LdapConfigParams{
		Context: ctx,
	}
}

// NewLdapConfigParamsWithHTTPClient creates a new LdapConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapConfigParamsWithHTTPClient(client *http.Client) *LdapConfigParams {
	return &LdapConfigParams{
		HTTPClient: client,
	}
}

/* LdapConfigParams contains all the parameters to send to the API endpoint
   for the ldap config operation.

   Typically these are written to a http.Request.
*/
type LdapConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapConfigParams) WithDefaults() *LdapConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap config params
func (o *LdapConfigParams) WithTimeout(timeout time.Duration) *LdapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap config params
func (o *LdapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap config params
func (o *LdapConfigParams) WithContext(ctx context.Context) *LdapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap config params
func (o *LdapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap config params
func (o *LdapConfigParams) WithHTTPClient(client *http.Client) *LdapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap config params
func (o *LdapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LdapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
