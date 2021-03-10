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

// NewDeleteUserCredentialsSamlParams creates a new DeleteUserCredentialsSamlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserCredentialsSamlParams() *DeleteUserCredentialsSamlParams {
	return &DeleteUserCredentialsSamlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsSamlParamsWithTimeout creates a new DeleteUserCredentialsSamlParams object
// with the ability to set a timeout on a request.
func NewDeleteUserCredentialsSamlParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsSamlParams {
	return &DeleteUserCredentialsSamlParams{
		timeout: timeout,
	}
}

// NewDeleteUserCredentialsSamlParamsWithContext creates a new DeleteUserCredentialsSamlParams object
// with the ability to set a context for a request.
func NewDeleteUserCredentialsSamlParamsWithContext(ctx context.Context) *DeleteUserCredentialsSamlParams {
	return &DeleteUserCredentialsSamlParams{
		Context: ctx,
	}
}

// NewDeleteUserCredentialsSamlParamsWithHTTPClient creates a new DeleteUserCredentialsSamlParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserCredentialsSamlParamsWithHTTPClient(client *http.Client) *DeleteUserCredentialsSamlParams {
	return &DeleteUserCredentialsSamlParams{
		HTTPClient: client,
	}
}

/* DeleteUserCredentialsSamlParams contains all the parameters to send to the API endpoint
   for the delete user credentials saml operation.

   Typically these are written to a http.Request.
*/
type DeleteUserCredentialsSamlParams struct {

	/* UserID.

	   id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user credentials saml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserCredentialsSamlParams) WithDefaults() *DeleteUserCredentialsSamlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user credentials saml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserCredentialsSamlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) WithTimeout(timeout time.Duration) *DeleteUserCredentialsSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) WithContext(ctx context.Context) *DeleteUserCredentialsSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) WithHTTPClient(client *http.Client) *DeleteUserCredentialsSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) WithUserID(userID int64) *DeleteUserCredentialsSamlParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user credentials saml params
func (o *DeleteUserCredentialsSamlParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
