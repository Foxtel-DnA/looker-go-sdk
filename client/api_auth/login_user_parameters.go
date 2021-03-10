// Code generated by go-swagger; DO NOT EDIT.

package api_auth

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

// NewLoginUserParams creates a new LoginUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoginUserParams() *LoginUserParams {
	return &LoginUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoginUserParamsWithTimeout creates a new LoginUserParams object
// with the ability to set a timeout on a request.
func NewLoginUserParamsWithTimeout(timeout time.Duration) *LoginUserParams {
	return &LoginUserParams{
		timeout: timeout,
	}
}

// NewLoginUserParamsWithContext creates a new LoginUserParams object
// with the ability to set a context for a request.
func NewLoginUserParamsWithContext(ctx context.Context) *LoginUserParams {
	return &LoginUserParams{
		Context: ctx,
	}
}

// NewLoginUserParamsWithHTTPClient creates a new LoginUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoginUserParamsWithHTTPClient(client *http.Client) *LoginUserParams {
	return &LoginUserParams{
		HTTPClient: client,
	}
}

/* LoginUserParams contains all the parameters to send to the API endpoint
   for the login user operation.

   Typically these are written to a http.Request.
*/
type LoginUserParams struct {

	/* Associative.

	   When true (default), API calls using the returned access_token are attributed to the admin user who created the access_token. When false, API activity is attributed to the user the access_token runs as. False requires a looker license.
	*/
	Associative *bool

	/* UserID.

	   Id of user.

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the login user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginUserParams) WithDefaults() *LoginUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the login user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the login user params
func (o *LoginUserParams) WithTimeout(timeout time.Duration) *LoginUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login user params
func (o *LoginUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login user params
func (o *LoginUserParams) WithContext(ctx context.Context) *LoginUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login user params
func (o *LoginUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login user params
func (o *LoginUserParams) WithHTTPClient(client *http.Client) *LoginUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login user params
func (o *LoginUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociative adds the associative to the login user params
func (o *LoginUserParams) WithAssociative(associative *bool) *LoginUserParams {
	o.SetAssociative(associative)
	return o
}

// SetAssociative adds the associative to the login user params
func (o *LoginUserParams) SetAssociative(associative *bool) {
	o.Associative = associative
}

// WithUserID adds the userID to the login user params
func (o *LoginUserParams) WithUserID(userID int64) *LoginUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the login user params
func (o *LoginUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *LoginUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Associative != nil {

		// query param associative
		var qrAssociative bool

		if o.Associative != nil {
			qrAssociative = *o.Associative
		}
		qAssociative := swag.FormatBool(qrAssociative)
		if qAssociative != "" {

			if err := r.SetQueryParam("associative", qAssociative); err != nil {
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
