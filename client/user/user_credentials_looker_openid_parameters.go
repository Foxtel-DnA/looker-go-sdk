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
	"github.com/go-openapi/swag"
)

// NewUserCredentialsLookerOpenidParams creates a new UserCredentialsLookerOpenidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCredentialsLookerOpenidParams() *UserCredentialsLookerOpenidParams {
	return &UserCredentialsLookerOpenidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCredentialsLookerOpenidParamsWithTimeout creates a new UserCredentialsLookerOpenidParams object
// with the ability to set a timeout on a request.
func NewUserCredentialsLookerOpenidParamsWithTimeout(timeout time.Duration) *UserCredentialsLookerOpenidParams {
	return &UserCredentialsLookerOpenidParams{
		timeout: timeout,
	}
}

// NewUserCredentialsLookerOpenidParamsWithContext creates a new UserCredentialsLookerOpenidParams object
// with the ability to set a context for a request.
func NewUserCredentialsLookerOpenidParamsWithContext(ctx context.Context) *UserCredentialsLookerOpenidParams {
	return &UserCredentialsLookerOpenidParams{
		Context: ctx,
	}
}

// NewUserCredentialsLookerOpenidParamsWithHTTPClient creates a new UserCredentialsLookerOpenidParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCredentialsLookerOpenidParamsWithHTTPClient(client *http.Client) *UserCredentialsLookerOpenidParams {
	return &UserCredentialsLookerOpenidParams{
		HTTPClient: client,
	}
}

/* UserCredentialsLookerOpenidParams contains all the parameters to send to the API endpoint
   for the user credentials looker openid operation.

   Typically these are written to a http.Request.
*/
type UserCredentialsLookerOpenidParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* UserID.

	   id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user credentials looker openid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsLookerOpenidParams) WithDefaults() *UserCredentialsLookerOpenidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user credentials looker openid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsLookerOpenidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) WithTimeout(timeout time.Duration) *UserCredentialsLookerOpenidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) WithContext(ctx context.Context) *UserCredentialsLookerOpenidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) WithHTTPClient(client *http.Client) *UserCredentialsLookerOpenidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) WithFields(fields *string) *UserCredentialsLookerOpenidParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) WithUserID(userID int64) *UserCredentialsLookerOpenidParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user credentials looker openid params
func (o *UserCredentialsLookerOpenidParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserCredentialsLookerOpenidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
