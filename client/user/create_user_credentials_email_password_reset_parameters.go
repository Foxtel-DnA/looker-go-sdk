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

// NewCreateUserCredentialsEmailPasswordResetParams creates a new CreateUserCredentialsEmailPasswordResetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserCredentialsEmailPasswordResetParams() *CreateUserCredentialsEmailPasswordResetParams {
	return &CreateUserCredentialsEmailPasswordResetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsEmailPasswordResetParamsWithTimeout creates a new CreateUserCredentialsEmailPasswordResetParams object
// with the ability to set a timeout on a request.
func NewCreateUserCredentialsEmailPasswordResetParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsEmailPasswordResetParams {
	return &CreateUserCredentialsEmailPasswordResetParams{
		timeout: timeout,
	}
}

// NewCreateUserCredentialsEmailPasswordResetParamsWithContext creates a new CreateUserCredentialsEmailPasswordResetParams object
// with the ability to set a context for a request.
func NewCreateUserCredentialsEmailPasswordResetParamsWithContext(ctx context.Context) *CreateUserCredentialsEmailPasswordResetParams {
	return &CreateUserCredentialsEmailPasswordResetParams{
		Context: ctx,
	}
}

// NewCreateUserCredentialsEmailPasswordResetParamsWithHTTPClient creates a new CreateUserCredentialsEmailPasswordResetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserCredentialsEmailPasswordResetParamsWithHTTPClient(client *http.Client) *CreateUserCredentialsEmailPasswordResetParams {
	return &CreateUserCredentialsEmailPasswordResetParams{
		HTTPClient: client,
	}
}

/* CreateUserCredentialsEmailPasswordResetParams contains all the parameters to send to the API endpoint
   for the create user credentials email password reset operation.

   Typically these are written to a http.Request.
*/
type CreateUserCredentialsEmailPasswordResetParams struct {

	/* Expires.

	   Expiring token.
	*/
	Expires *bool

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

// WithDefaults hydrates default values in the create user credentials email password reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserCredentialsEmailPasswordResetParams) WithDefaults() *CreateUserCredentialsEmailPasswordResetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user credentials email password reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserCredentialsEmailPasswordResetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithTimeout(timeout time.Duration) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithContext(ctx context.Context) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithHTTPClient(client *http.Client) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpires adds the expires to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithExpires(expires *bool) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetExpires(expires)
	return o
}

// SetExpires adds the expires to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetExpires(expires *bool) {
	o.Expires = expires
}

// WithFields adds the fields to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithFields(fields *string) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) WithUserID(userID int64) *CreateUserCredentialsEmailPasswordResetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create user credentials email password reset params
func (o *CreateUserCredentialsEmailPasswordResetParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsEmailPasswordResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expires != nil {

		// query param expires
		var qrExpires bool

		if o.Expires != nil {
			qrExpires = *o.Expires
		}
		qExpires := swag.FormatBool(qrExpires)
		if qExpires != "" {

			if err := r.SetQueryParam("expires", qExpires); err != nil {
				return err
			}
		}
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
