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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewUpdateLdapConfigParams creates a new UpdateLdapConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLdapConfigParams() *UpdateLdapConfigParams {
	return &UpdateLdapConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapConfigParamsWithTimeout creates a new UpdateLdapConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateLdapConfigParamsWithTimeout(timeout time.Duration) *UpdateLdapConfigParams {
	return &UpdateLdapConfigParams{
		timeout: timeout,
	}
}

// NewUpdateLdapConfigParamsWithContext creates a new UpdateLdapConfigParams object
// with the ability to set a context for a request.
func NewUpdateLdapConfigParamsWithContext(ctx context.Context) *UpdateLdapConfigParams {
	return &UpdateLdapConfigParams{
		Context: ctx,
	}
}

// NewUpdateLdapConfigParamsWithHTTPClient creates a new UpdateLdapConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLdapConfigParamsWithHTTPClient(client *http.Client) *UpdateLdapConfigParams {
	return &UpdateLdapConfigParams{
		HTTPClient: client,
	}
}

/* UpdateLdapConfigParams contains all the parameters to send to the API endpoint
   for the update ldap config operation.

   Typically these are written to a http.Request.
*/
type UpdateLdapConfigParams struct {

	/* Body.

	   LDAP Config
	*/
	Body *models.LDAPConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update ldap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapConfigParams) WithDefaults() *UpdateLdapConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update ldap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update ldap config params
func (o *UpdateLdapConfigParams) WithTimeout(timeout time.Duration) *UpdateLdapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap config params
func (o *UpdateLdapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap config params
func (o *UpdateLdapConfigParams) WithContext(ctx context.Context) *UpdateLdapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap config params
func (o *UpdateLdapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap config params
func (o *UpdateLdapConfigParams) WithHTTPClient(client *http.Client) *UpdateLdapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap config params
func (o *UpdateLdapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update ldap config params
func (o *UpdateLdapConfigParams) WithBody(body *models.LDAPConfig) *UpdateLdapConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update ldap config params
func (o *UpdateLdapConfigParams) SetBody(body *models.LDAPConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
