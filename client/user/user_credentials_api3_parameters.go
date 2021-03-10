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

// NewUserCredentialsApi3Params creates a new UserCredentialsApi3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCredentialsApi3Params() *UserCredentialsApi3Params {
	return &UserCredentialsApi3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCredentialsApi3ParamsWithTimeout creates a new UserCredentialsApi3Params object
// with the ability to set a timeout on a request.
func NewUserCredentialsApi3ParamsWithTimeout(timeout time.Duration) *UserCredentialsApi3Params {
	return &UserCredentialsApi3Params{
		timeout: timeout,
	}
}

// NewUserCredentialsApi3ParamsWithContext creates a new UserCredentialsApi3Params object
// with the ability to set a context for a request.
func NewUserCredentialsApi3ParamsWithContext(ctx context.Context) *UserCredentialsApi3Params {
	return &UserCredentialsApi3Params{
		Context: ctx,
	}
}

// NewUserCredentialsApi3ParamsWithHTTPClient creates a new UserCredentialsApi3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUserCredentialsApi3ParamsWithHTTPClient(client *http.Client) *UserCredentialsApi3Params {
	return &UserCredentialsApi3Params{
		HTTPClient: client,
	}
}

/* UserCredentialsApi3Params contains all the parameters to send to the API endpoint
   for the user credentials api3 operation.

   Typically these are written to a http.Request.
*/
type UserCredentialsApi3Params struct {

	/* CredentialsApi3ID.

	   Id of API 3 Credential

	   Format: int64
	*/
	CredentialsApi3ID int64

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* UserID.

	   Id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user credentials api3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsApi3Params) WithDefaults() *UserCredentialsApi3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user credentials api3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsApi3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithTimeout(timeout time.Duration) *UserCredentialsApi3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithContext(ctx context.Context) *UserCredentialsApi3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithHTTPClient(client *http.Client) *UserCredentialsApi3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialsApi3ID adds the credentialsApi3ID to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithCredentialsApi3ID(credentialsApi3ID int64) *UserCredentialsApi3Params {
	o.SetCredentialsApi3ID(credentialsApi3ID)
	return o
}

// SetCredentialsApi3ID adds the credentialsApi3Id to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetCredentialsApi3ID(credentialsApi3ID int64) {
	o.CredentialsApi3ID = credentialsApi3ID
}

// WithFields adds the fields to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithFields(fields *string) *UserCredentialsApi3Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the user credentials api3 params
func (o *UserCredentialsApi3Params) WithUserID(userID int64) *UserCredentialsApi3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user credentials api3 params
func (o *UserCredentialsApi3Params) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserCredentialsApi3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credentials_api3_id
	if err := r.SetPathParam("credentials_api3_id", swag.FormatInt64(o.CredentialsApi3ID)); err != nil {
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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
