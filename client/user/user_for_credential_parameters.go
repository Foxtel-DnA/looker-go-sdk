// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserForCredentialParams creates a new UserForCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserForCredentialParams() *UserForCredentialParams {
	return &UserForCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserForCredentialParamsWithTimeout creates a new UserForCredentialParams object
// with the ability to set a timeout on a request.
func NewUserForCredentialParamsWithTimeout(timeout time.Duration) *UserForCredentialParams {
	return &UserForCredentialParams{
		timeout: timeout,
	}
}

// NewUserForCredentialParamsWithContext creates a new UserForCredentialParams object
// with the ability to set a context for a request.
func NewUserForCredentialParamsWithContext(ctx context.Context) *UserForCredentialParams {
	return &UserForCredentialParams{
		Context: ctx,
	}
}

// NewUserForCredentialParamsWithHTTPClient creates a new UserForCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserForCredentialParamsWithHTTPClient(client *http.Client) *UserForCredentialParams {
	return &UserForCredentialParams{
		HTTPClient: client,
	}
}

/* UserForCredentialParams contains all the parameters to send to the API endpoint
   for the user for credential operation.

   Typically these are written to a http.Request.
*/
type UserForCredentialParams struct {

	/* CredentialID.

	   Id of credential
	*/
	CredentialID string

	/* CredentialType.

	   Type name of credential
	*/
	CredentialType string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user for credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserForCredentialParams) WithDefaults() *UserForCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user for credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserForCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user for credential params
func (o *UserForCredentialParams) WithTimeout(timeout time.Duration) *UserForCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user for credential params
func (o *UserForCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user for credential params
func (o *UserForCredentialParams) WithContext(ctx context.Context) *UserForCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user for credential params
func (o *UserForCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user for credential params
func (o *UserForCredentialParams) WithHTTPClient(client *http.Client) *UserForCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user for credential params
func (o *UserForCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the user for credential params
func (o *UserForCredentialParams) WithCredentialID(credentialID string) *UserForCredentialParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the user for credential params
func (o *UserForCredentialParams) SetCredentialID(credentialID string) {
	o.CredentialID = credentialID
}

// WithCredentialType adds the credentialType to the user for credential params
func (o *UserForCredentialParams) WithCredentialType(credentialType string) *UserForCredentialParams {
	o.SetCredentialType(credentialType)
	return o
}

// SetCredentialType adds the credentialType to the user for credential params
func (o *UserForCredentialParams) SetCredentialType(credentialType string) {
	o.CredentialType = credentialType
}

// WithFields adds the fields to the user for credential params
func (o *UserForCredentialParams) WithFields(fields *string) *UserForCredentialParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user for credential params
func (o *UserForCredentialParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *UserForCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credential_id
	if err := r.SetPathParam("credential_id", o.CredentialID); err != nil {
		return err
	}

	// path param credential_type
	if err := r.SetPathParam("credential_type", o.CredentialType); err != nil {
		return err
	}

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
