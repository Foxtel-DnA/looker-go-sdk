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

// NewUserCredentialsEmailParams creates a new UserCredentialsEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCredentialsEmailParams() *UserCredentialsEmailParams {
	return &UserCredentialsEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCredentialsEmailParamsWithTimeout creates a new UserCredentialsEmailParams object
// with the ability to set a timeout on a request.
func NewUserCredentialsEmailParamsWithTimeout(timeout time.Duration) *UserCredentialsEmailParams {
	return &UserCredentialsEmailParams{
		timeout: timeout,
	}
}

// NewUserCredentialsEmailParamsWithContext creates a new UserCredentialsEmailParams object
// with the ability to set a context for a request.
func NewUserCredentialsEmailParamsWithContext(ctx context.Context) *UserCredentialsEmailParams {
	return &UserCredentialsEmailParams{
		Context: ctx,
	}
}

// NewUserCredentialsEmailParamsWithHTTPClient creates a new UserCredentialsEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCredentialsEmailParamsWithHTTPClient(client *http.Client) *UserCredentialsEmailParams {
	return &UserCredentialsEmailParams{
		HTTPClient: client,
	}
}

/* UserCredentialsEmailParams contains all the parameters to send to the API endpoint
   for the user credentials email operation.

   Typically these are written to a http.Request.
*/
type UserCredentialsEmailParams struct {

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

// WithDefaults hydrates default values in the user credentials email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsEmailParams) WithDefaults() *UserCredentialsEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user credentials email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCredentialsEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user credentials email params
func (o *UserCredentialsEmailParams) WithTimeout(timeout time.Duration) *UserCredentialsEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user credentials email params
func (o *UserCredentialsEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user credentials email params
func (o *UserCredentialsEmailParams) WithContext(ctx context.Context) *UserCredentialsEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user credentials email params
func (o *UserCredentialsEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user credentials email params
func (o *UserCredentialsEmailParams) WithHTTPClient(client *http.Client) *UserCredentialsEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user credentials email params
func (o *UserCredentialsEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the user credentials email params
func (o *UserCredentialsEmailParams) WithFields(fields *string) *UserCredentialsEmailParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user credentials email params
func (o *UserCredentialsEmailParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the user credentials email params
func (o *UserCredentialsEmailParams) WithUserID(userID int64) *UserCredentialsEmailParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user credentials email params
func (o *UserCredentialsEmailParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserCredentialsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
