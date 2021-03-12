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

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// NewCreateUserCredentialsApi3Params creates a new CreateUserCredentialsApi3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserCredentialsApi3Params() *CreateUserCredentialsApi3Params {
	return &CreateUserCredentialsApi3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserCredentialsApi3ParamsWithTimeout creates a new CreateUserCredentialsApi3Params object
// with the ability to set a timeout on a request.
func NewCreateUserCredentialsApi3ParamsWithTimeout(timeout time.Duration) *CreateUserCredentialsApi3Params {
	return &CreateUserCredentialsApi3Params{
		timeout: timeout,
	}
}

// NewCreateUserCredentialsApi3ParamsWithContext creates a new CreateUserCredentialsApi3Params object
// with the ability to set a context for a request.
func NewCreateUserCredentialsApi3ParamsWithContext(ctx context.Context) *CreateUserCredentialsApi3Params {
	return &CreateUserCredentialsApi3Params{
		Context: ctx,
	}
}

// NewCreateUserCredentialsApi3ParamsWithHTTPClient creates a new CreateUserCredentialsApi3Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserCredentialsApi3ParamsWithHTTPClient(client *http.Client) *CreateUserCredentialsApi3Params {
	return &CreateUserCredentialsApi3Params{
		HTTPClient: client,
	}
}

/* CreateUserCredentialsApi3Params contains all the parameters to send to the API endpoint
   for the create user credentials api3 operation.

   Typically these are written to a http.Request.
*/
type CreateUserCredentialsApi3Params struct {

	/* Body.

	   API 3 Credential
	*/
	Body *models.CredentialsApi3

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

// WithDefaults hydrates default values in the create user credentials api3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserCredentialsApi3Params) WithDefaults() *CreateUserCredentialsApi3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user credentials api3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserCredentialsApi3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithTimeout(timeout time.Duration) *CreateUserCredentialsApi3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithContext(ctx context.Context) *CreateUserCredentialsApi3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithHTTPClient(client *http.Client) *CreateUserCredentialsApi3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithBody(body *models.CredentialsApi3) *CreateUserCredentialsApi3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetBody(body *models.CredentialsApi3) {
	o.Body = body
}

// WithFields adds the fields to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithFields(fields *string) *CreateUserCredentialsApi3Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) WithUserID(userID int64) *CreateUserCredentialsApi3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create user credentials api3 params
func (o *CreateUserCredentialsApi3Params) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserCredentialsApi3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
